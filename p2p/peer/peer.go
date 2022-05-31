package peer

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"

	pb "serkanmulayim/blockchain/p2p/proto"
	"serkanmulayim/blockchain/storage"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	gpeer "google.golang.org/grpc/peer"
)

const (
	VERSION = "0.0.1"

	STATUS_SUCCESS = "success"
	STATUS_REJECT  = "reject"
	STATUS_REQUEST = "request"
)

var (
	//myAddrs               = flag.String("myaddrs", "10.0.0.3:31000,10.0.0.254:31000,10.0.0.175:31000", "my ips") // in order to avoid connections to self
	port                = flag.Int("port", 31000, "server port")
	peers               = flag.String("peers", "10.0.0.175:31000,10.0.0.254:31000", "initial peers to connect, (comma separated)")
	maxOnlinePoolCount  = flag.Int("onlinePeers", 4, "max number of peers in the connection pool")
	discardInitialPeers = flag.Bool("discardInitialPeers", false, "switch to discard nitial peers list")

	poolOnlineMu sync.RWMutex
	onlinePool   map[string]onlinePoolObject = make(map[string]onlinePoolObject)

	onlinePoolMonitorWaitPeriod time.Duration = 30 //sec
	getPeersWaitPeriod          time.Duration = 22 //sec

	initialAddrs []string
	startOnce    sync.Once
)

type P2PServer struct {
	pb.UnimplementedP2PServer
}

type P2PClient struct {
	cc grpc.ClientConnInterface
}

type onlinePoolEntry struct {
	addr string
	obj  onlinePoolObject
}

type onlinePoolObject struct {
	conn   *grpc.ClientConn
	client pb.P2PClient
}

//TODO: understand GRPC server limits (associate it with maxOnlinePoolCount), DDoS attacks
//      https://www.evanjones.ca/rate-limit-for-reliable-servers.html

//online pool monitoring worker for maximizing the online pool connections in limits
func onlinePoolMonitor() {
	for {
		time.Sleep(onlinePoolMonitorWaitPeriod * time.Second)
		updatePool()
	}
}

func getPeersJob() {
	for {
		time.Sleep(getPeersWaitPeriod * time.Second)
		doGetPeersRequests()
	}
}

func updatePool() {
	log.Debugf("UpdatePool function is called")
	//clean up non-working connections
	poolOnlineMu.Lock()
	for key, obj := range onlinePool {
		state := obj.conn.GetState()
		log.Debugf("State for %v is %v", key, state)
		if state != connectivity.Ready && state != connectivity.Connecting {
			log.Infof("Node %v is removed from online pool since its state is %v", key, state)
			obj.conn.Close()

			delete(onlinePool, key)
		}
	}
	size := len(onlinePool)
	poolOnlineMu.Unlock()

	connectToPeers(size)
}

//Server implementation of SendHello in P2P protocol
func (s *P2PServer) SendHello(ctx context.Context, in *pb.HelloMessage) (*pb.HelloMessage, error) {
	if in.Status == STATUS_REQUEST {
		p, _ := gpeer.FromContext(ctx)
		log.Debugf("Received HELLO %v from %v, %v\n", in, p.Addr.String(), p.Addr.Network())
		addr := fmt.Sprintf("%v:%d", getIp(p), in.Port)

		defer func() {
			storage.PeersReceiveChannel <- storage.OfflinePoolObject{Addr: addr, TryCount: 0}
		}()
	}
	return &pb.HelloMessage{Agent: "Agent", Version: VERSION, Port: uint32(*port), Status: STATUS_SUCCESS}, nil
}

func (s *P2PServer) GetPeers(ctx context.Context, req *pb.PeersRequest) (*pb.PeersResponse, error) {
	p, _ := gpeer.FromContext(ctx)
	log.Debugf("Received PEERS %v from %v, %v\n", req, p.Addr.String(), p.Addr.Network())
	storage.PeerArrayMu.RLock()
	defer storage.PeerArrayMu.RUnlock()
	return &pb.PeersResponse{Peers: storage.PeerArray, Status: STATUS_SUCCESS}, nil
}

// func (s *P2PServer) GetObject(ctx context.Context, objId *pb.ObjectId) (*pb.Object, error) {
// 	obj, err := storage.GetObject(objId.ObjectId)

// 	return nil, nil
// }
// func (s *P2PServer) IHaveObject(ctx context.Context, objId *pb.ObjectId) (*pb.Empty, error) {
// 	storage.HasObject(objId.ObjectId)
// }

func Start() error {
	startOnce.Do(start)

	return nil
}

func Stop() {
	storage.StopLevelDb()
	closeConnections()
}

func closeConnections() {
	log.Info("Closing connections...")
	poolOnlineMu.Lock()
	defer poolOnlineMu.Unlock()
	for _, obj := range onlinePool {
		obj.conn.Close()
	}
	log.Info("...Done")
}

func start() {

	flag.Parse()

	err := initPools()
	if err != nil {
		log.Fatalf("cannot init pools: %v", err)
	}

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

		if err != nil {
			log.Fatalf("error in listening port: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterP2PServer(s, &P2PServer{})
		log.Infof("Started server at %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve the P2P server: %v", err)
		}
	}()

	if err := connectToPeers(0); err != nil {
		log.Fatalf("Issue with connecting to peers", err)
	}

	doGetPeersRequests()

	go onlinePoolMonitor()
	go getPeersJob()
}

func connectToPeers(size int) error {
	var wg sync.WaitGroup
	stopCh := make(chan bool)
	workerCh := make(chan *storage.OfflinePoolObject)
	var atomicCounter int32 = 0

	for i := int(size); i < *maxOnlinePoolCount; i++ {
		wg.Add(1)
		go sendHello(workerCh, &wg, STATUS_REQUEST, &atomicCounter, stopCh)
	}

	for obj := range storage.PeerGenerator(initialAddrs, *discardInitialPeers, stopCh) {
		if atomic.LoadInt32(&atomicCounter) >= int32(*maxOnlinePoolCount-size) {
			stopCh <- true
		} else {
			if obj.Addr == "" {
				log.Debugf("DUUUUUUUT %v", obj)
			}
			obj2 := obj
			workerCh <- &obj2
		}
	}
	close(workerCh)
	wg.Wait()

	poolOnlineMu.RLock()
	poolSize := len(onlinePool)
	poolOnlineMu.RUnlock()

	log.Infof("Connected to %d peers with the initial connect to peers request.", poolSize)
	return nil
}

//tries to send one successful hello message
func sendHello(ch chan *storage.OfflinePoolObject, wg *sync.WaitGroup, status string, atomicCounter *int32, stopChannel <-chan bool) {
	if wg != nil {
		defer wg.Done()
	}
	for {
		poolOnlineMu.RLock()
		if len(onlinePool) >= *maxOnlinePoolCount {
			poolOnlineMu.RUnlock()
			return
		}

		var obj *storage.OfflinePoolObject
		select {
		case <-stopChannel: // if online pool counter is high
			poolOnlineMu.RUnlock()
			return
		case obj = <-ch:
		}

		if obj == nil { //close workerCh (if offline pool is small)
			poolOnlineMu.RUnlock()
			return
		}
		_, ok := onlinePool[obj.Addr]
		if ok {
			poolOnlineMu.RUnlock()
			continue
		}
		poolOnlineMu.RUnlock()

		conn, err := grpc.Dial(obj.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Infof("Could not connect to %v, %v", obj.Addr, err)
			storage.PeersReceiveChannel <- storage.OfflinePoolObject{TryCount: obj.TryCount + 1, Addr: obj.Addr}
			continue
		}
		conn.Target()

		c := pb.NewP2PClient(conn)

		// ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := c.SendHello(ctx, &pb.HelloMessage{Agent: "xxx", Version: VERSION, Port: uint32(*port), Status: status})
		if err == nil {
			if res.Status != STATUS_SUCCESS {
				log.Debugf("Skipping unsuccessful hello message")
				continue
			}
			storage.PeersReceiveChannel <- storage.OfflinePoolObject{TryCount: 0, Addr: obj.Addr}
			onlinePoolObj := onlinePoolObject{client: c, conn: conn}
			entry := onlinePoolEntry{obj: onlinePoolObj, addr: obj.Addr}
			poolOnlineMu.Lock()
			defer poolOnlineMu.Unlock()
			if len(onlinePool) < *maxOnlinePoolCount {

				onlinePool[entry.addr] = entry.obj

				log.Infof("Online Pool Updated with %v.", entry.addr)
			} else {
				log.Infof("Online Pool CANNOT be updated with %v since it is maxed out. Skipping this one...", entry.addr)
			}
			return

		} else {
			log.Debugf("Error in SendHello for %v :%v", obj.Addr, err)
			storage.PeersReceiveChannel <- storage.OfflinePoolObject{TryCount: obj.TryCount + 1, Addr: obj.Addr}
		}
	}
}

func doGetPeersRequests() {
	var wg sync.WaitGroup
	var newPeersMu sync.Mutex
	var newPeersSet map[string]bool = make(map[string]bool)
	var newPeersArray []string

	poolOnlineMu.RLock()
	defer poolOnlineMu.RUnlock()
	for key, obj := range onlinePool {
		wg.Add(1)
		go getPeers(key, &obj, &wg, newPeersSet, &newPeersMu)
	}
	wg.Wait()
	log.Info("Finished with Peers requests")
	newPeersArray = make([]string, 0, len(newPeersSet))
	for peer := range newPeersSet {
		newPeersArray = append(newPeersArray, peer)
	}
	log.Debug(newPeersArray)

	storage.PeersBulkReceiveChannel <- newPeersArray
}

//not thread safe for online pool
func getPeers(addr string, obj *onlinePoolObject, wg *sync.WaitGroup, newPeers map[string]bool, newPeersMu *sync.Mutex) {
	log.Debugf("Sent PEER request to %v", addr)
	defer wg.Done()
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	res, err := obj.client.GetPeers(ctx, &pb.PeersRequest{})
	if err != nil {
		log.Infof("Could not get peers from %v, %v. Connection will eventually be closed.", addr, err)
		return
	}

	if res.Status != STATUS_SUCCESS {
		log.Infof("Get peers request for %v did not return SUCCESS, but returned %v", addr, res.Status)
	}

	newPeersMu.Lock()
	for _, peer := range res.Peers {
		newPeers[peer] = true
	}
	newPeersMu.Unlock()

}

func initPools() (err error) {
	initialAddrs = strings.Split(*peers, ",")
	storage.StartLevelDB(*port)

	err = storage.StartOfflinePeersPool(*port, initialAddrs)
	return
}

func getIp(peer *gpeer.Peer) string {

	ipport := peer.Addr.String()
	var i int
	for i = len(ipport) - 1; i > 0; i-- {
		if ipport[i] == ':' {
			break
		}
	}
	//peer will have the string always in the correct format, so this is not a case
	return string(ipport[:i])
}

// func getRandomB64String(numBytes int) string {
// 	bytes := make([]byte, numBytes)
// 	rand.Read(bytes)
// 	return base64.URLEncoding.EncodeToString(bytes)
// }
