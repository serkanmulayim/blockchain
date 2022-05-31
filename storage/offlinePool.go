package storage

import (
	"errors"
	"net"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	//pool states
	Stopped      int8 = 0
	Initializing      = 1
	Running           = 2

	maxTryCount   uint8 = 4
	maxObjects    int   = 16384
	levelDBPrefix       = "peers"

	peerArrayUpdateInterval time.Duration = 10 //seconds

)

type OfflinePoolObject struct {
	TryCount uint8
	Addr     string
}

var (
	offlinePoolMu sync.RWMutex
	offlinePool   map[string]OfflinePoolObject = make(map[string]OfflinePoolObject)
	state         int8

	PeersReceiveChannel     chan OfflinePoolObject = make(chan OfflinePoolObject, 100)
	PeersBulkReceiveChannel chan []string          = make(chan []string)

	PeerArrayMu sync.RWMutex
	PeerArray   []string
)

//including port for ipv4 and ipv6
func verifyAddr(ipport string) bool {
	var i int
	for i = len(ipport) - 1; i > 0; i-- {
		if ipport[i] == ':' {
			break
		}
	}
	if i == 0 || i == len(ipport)-1 {
		return false
	}
	port := ipport[i+1:]
	portI, err := strconv.Atoi(port)
	if err != nil || portI < 0 || portI > 65535 {
		return false
	}

	ip := net.ParseIP(ipport[:i])
	return ip != nil
}

func dbWrite(obj OfflinePoolObject) {
	if verifyAddr(obj.Addr) {
		PeersDB.Put([]byte(obj.Addr), []byte{obj.TryCount}, nil)
	}
}

func dbDelete(addr string) {
	PeersDB.Delete([]byte(addr), nil)
}

func initPool() error {
	offlinePoolMu.Lock()
	defer offlinePoolMu.Unlock()
	if state == Stopped {
		offlinePool = make(map[string]OfflinePoolObject)
		state = Initializing
		getOfflinePool()
		go startPeerListUpdater()
		state = Running

	} else if state == Initializing {
		return errors.New("Pool is already initializing")
	} else if state == Running {
		return errors.New("Pool is already running")
	} else {
		log.Fatalf("Unknown state for the pool :%d", state)
	}
	return nil
}

func startPeerListUpdater() {
	for {
		updatePeersList()
		time.Sleep(peerArrayUpdateInterval * time.Second)
	}
}

func updatePeersList() {
	offlinePoolMu.RLock()
	defer offlinePoolMu.RUnlock()
	PeerArrayMu.Lock()
	defer PeerArrayMu.Unlock()

	PeerArray = make([]string, 0, len(offlinePool))
	for k, v := range offlinePool {
		if v.TryCount == 0 {
			PeerArray = append(PeerArray, k)
		}
	}
}

//TODO: check if offlinePool cache/map is still necessary when using leveldb.
//      it might be necessary for a randomized strategy for iterating the peers
// not thread safe
func getOfflinePool() {

	iter := PeersDB.NewIterator(nil, nil)
	count := 0
	for iter.Next() {
		if count >= maxObjects {
			break
		}
		key := string(iter.Key())
		offlinePool[key] = OfflinePoolObject{Addr: key, TryCount: iter.Value()[0]}
		count++
		log.Debugf("%s %d", key, iter.Value()[0])
	}
	iter.Release()
}

//Manages the pool
func StartOfflinePeersPool(port int, initAddr []string) error {

	offlinePoolMu.RLock()
	if state != Stopped {
		offlinePoolMu.RUnlock()
		return errors.New("Offline Pool was already started ")
	}
	offlinePoolMu.RUnlock()

	err := initPool()
	if err != nil {
		log.Fatalf("Cannot init the connection pool with reason: %v", err)
	}
	go offlinePoolUpdater(initAddr)
	return nil
}

func offlinePoolUpdater(initAddr []string) {
	for {
		select {

		case obj := <-PeersReceiveChannel:
			offlinePoolMu.Lock()
			if obj.TryCount >= maxTryCount {
				//delete
				delete(offlinePool, obj.Addr)
				offlinePoolMu.Unlock()
				dbDelete(obj.Addr)
				log.Debugf("%v is removed from the offline peer pool", obj.Addr)
			} else {
				//update db

				if len(offlinePool) >= maxObjects {
					offlinePoolMu.Unlock()
					log.Debugf("%v cannot be added to the offline peer pool, since the pool is full. Discarding obj...", obj.Addr)
				} else {
					if !arrayContains(initAddr, obj.Addr) { //do not put initial peers to db
						offlinePool[obj.Addr] = obj
						offlinePoolMu.Unlock()
						dbWrite(obj)
						log.Debugf("%v is added to the offline peer pool with try count:%d...", obj.Addr, obj.TryCount)
					} else {
						offlinePoolMu.Unlock()
						log.Debugf("Not putting %v to offline pool since it is one of the initial peers.", obj.Addr)
					}
				}
			}
		case bulk := <-PeersBulkReceiveChannel:
			offlinePoolMu.Lock()
			for _, peer := range bulk {
				if len(offlinePool) >= maxObjects {
					break
				}
				_, ok := offlinePool[peer]
				if !ok {
					exist, err := PeersDB.Has([]byte(peer), nil)
					if err != nil {
						log.Errorf("DB Has operation failing with:%v", err)
					}
					if !exist {
						obj := OfflinePoolObject{Addr: peer, TryCount: 0}
						dbWrite(obj)
						offlinePool[peer] = obj
					}
				}
			}
			offlinePoolMu.Unlock()
		}
	}
}

func arrayContains(a []string, addr string) bool {
	for _, s := range a {
		if addr == s {
			return true
		}
	}
	return false
}

func PeerGenerator(initialAddrs []string, discardInitialAddr bool, stopChannel chan bool) <-chan OfflinePoolObject {
	offlinePoolMu.RLock()
	defer offlinePoolMu.RUnlock()

	if state != Running {
		log.Panicf("Offline Pool has not started yet")
	}
	ch := make(chan OfflinePoolObject)

	go func() {
		for key, val := range offlinePool {
			val.Addr = key
			select {
			case ch <- val:

			case <-stopChannel:
				close(ch)
				return
			}
		}
		if !discardInitialAddr {
			for _, addr := range initialAddrs {
				if _, ok := offlinePool[addr]; !ok {

					out := OfflinePoolObject{Addr: addr, TryCount: 0}
					select {
					case ch <- out:

					case <-stopChannel:
						close(ch)
						return
					}
				}
			}
		}

		close(ch)
	}()

	return ch
}
