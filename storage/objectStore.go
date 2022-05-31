package storage

import (
	"crypto/sha256"
	"encoding/json"
	"errors"

	jc "serkanmulayim/blockchain/imports/jsoncanonicalizer"
	pb "serkanmulayim/blockchain/p2p/proto"
)

const (
	ObjectTypeBlock = "block"
	ObjectTypeTx    = "tx"
)

var (
	ObjectReceiveCh chan string
)

func GetObject(objId []byte, typee string) ([]byte, error) {

	switch typee {
	case ObjectTypeBlock:
		return BlockObjectsDB.Get(objId, nil)
	case ObjectTypeTx:
		return TxObjectsDB.Get(objId, nil)
	default:
		return nil, errors.New("Unknown object type for GETOBJECT:" + typee)
	}
}

func HasObject(objId []byte, typee string) (bool, error) {
	switch typee {
	case ObjectTypeBlock:
		return BlockObjectsDB.Has(objId, nil)
	case ObjectTypeTx:
		return TxObjectsDB.Has(objId, nil)
	default:
		return false, errors.New("Unknown object type for HASOBJECT:" + typee)
	}

}

func PutObject(obj []byte, typee string) error {

	switch typee {
	case ObjectTypeBlock:
		var b pb.Block
		err := json.Unmarshal(obj, &b)
		if err != nil {
			return err
		}

		in, err := json.Marshal(b)
		//not expected
		if err != nil {
			return err
		}

		out, err := jc.Transform(in)
		if err != nil { //not expected
			return err
		}
		objId := sha256.Sum256(out)

		return BlockObjectsDB.Put(objId[:], out, nil)
	case ObjectTypeTx:
		var t pb.Tx
		err := json.Unmarshal(obj, &t)
		if err != nil {
			return err
		}
		in, err := json.Marshal(t)
		if err != nil { //not expected
			return err
		}

		out, err := jc.Transform(in)
		if err != nil { //not expected
			return err
		}
		objId := sha256.Sum256(out)

		return TxObjectsDB.Put(objId[:], obj, nil)
	default:
		return errors.New("Unknown object type for PUTOBJECT:" + typee)
	}

}
