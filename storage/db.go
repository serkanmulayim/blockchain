package storage

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	PeersPrefix        = "peers"
	BlockObjectsPrefix = "blocks"
	TxObjectsPrefix    = "txs"
)

var (
	PeersDB        *leveldb.DB = nil
	BlockObjectsDB *leveldb.DB = nil
	TxObjectsDB    *leveldb.DB = nil
)

func startLevelDB(prefix string, port int) (*leveldb.DB, error) { //port is used as a namespace for the separation between node data in local
	dbLoc := fmt.Sprintf("./db/%v.db", prefix)
	if port != 0 {
		dbLoc = fmt.Sprintf("./db/%d%v.db", port, prefix)
	}
	db, err := leveldb.OpenFile(dbLoc, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func StartLevelDB(port int) {
	db, err := startLevelDB(PeersPrefix, port)
	if err != nil {
		log.Fatalf("Could not initiate the PeersDB")
	}
	PeersDB = db

	db, err = startLevelDB(BlockObjectsPrefix, port)
	if err != nil {
		log.Fatalf("Could not initiate the BlockObjectsDB")
	}
	BlockObjectsDB = db

	db, err = startLevelDB(TxObjectsPrefix, port)
	if err != nil {
		log.Fatalf("Could not initiate the TxObjectsDB")
	}
	TxObjectsDB = db
}

func StopLevelDb() {
	log.Info("Shutting down DB...")
	PeersDB.Close()
	BlockObjectsDB.Close()
	TxObjectsDB.Close()
	log.Info("...Done")
}
