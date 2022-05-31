package storage

//NOT USED
//needed for serialization in json format for []byte types. Default serializers
//such as protojson and json do not HEX encode them, but to base64 encode. Not desirable

// import (
// 	"encoding/hex"
// 	"encoding/json"
// 	"errors"
// 	pb "serkanmulayim/blockchain/p2p/proto"

// 	jc "serkanmulayim/blockchain/imports/jsoncanonicalizer"

// 	"google.golang.org/protobuf/types/known/anypb"
// )

// const (
// 	ObjectTypeBlock = "blk"
// 	ObjectTypeTx    = "tx"
// )

// type Block struct {
// 	Created uint64   `json:"created,omitempty"`
// 	Miner   string   `json:"miner,omitempty"`
// 	Nonce   string   `json:"nonce,omitempty"`
// 	Note    string   `json:"note,omitempty"`
// 	PrevId  string   `json:"prevId,omitempty"`
// 	T       string   `json:"t,omitempty"`
// 	TxIds   []string `json:"txIds,omitempty"`
// }

// type Tx struct {
// 	Inputs  []*Input  `json:"inputs,omitempty"`
// 	Outputs []*Output `json:"outputs,omitempty"`
// 	Height  int64     `json:"height,omitempty"`
// }

// type OutPoint struct {
// 	Index int64  `json:"index,omitempty"`
// 	TxId  string `json:"txId,omitempty"`
// }

// type Input struct {
// 	Outpoint *OutPoint `json:"outpoint,omitempty"`
// 	Sig      []byte    `json:"sig"`
// }

// type Output struct {
// 	Pubkey string `json:"pubkey,omitempty"`
// 	Value  int64  `json:"value,omitempty"`
// }

// func ProtoToJson(obj *anypb.Any) ([]byte, error) {
// 	msg, err := obj.UnmarshalNew()
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch t := msg.(type) {
// 	case *pb.Block:
// 		return protoBlockToBlock(t)
// 	case *pb.Tx:
// 		return protoTxToTx(t)
// 	default:
// 		return nil, errors.New("Unknown type of block")
// 	}
// }

// func JsonToProto(in []byte) (*anypb.Any, error) {
// 	var dat map[string]interface{}
// 	if err := json.Unmarshal(in, &dat); err != nil {
// 		return nil, err
// 	}
// 	switch {
// 	case dat["type"].(string) == ObjectTypeBlock:
// 		return blkToProtoBlk(in)
// 	case dat["type"].(string) == ObjectTypeTx:
// 		return txToProtoTx(in)
// 	default:
// 		return nil, errors.New("Unknown Type for the object retrieved from DB")
// 	}
// }

// func txToProtoTx(in []byte) (*anypb.Any, error) {
// 	var pbtx pb.Tx
// 	var tx Tx

// 	if err := json.Unmarshal(in, &tx); err != nil {
// 		return nil, err
// 	}

// 	pinputs := make([]*pb.Input, 0, len(tx.Inputs))

// 	for _, input := range tx.Inputs {
// 		var pinput pb.Input
// 		pinput.Sig = input.Sig

// 		var poutpoint pb.OutPoint
// 		poutpoint.Index = input.Outpoint.Index

// 		if v, err := hex.DecodeString(input.Outpoint.TxId); err != nil {
// 			return nil, errors.New("Input outpoint txId is not hex decodable for the object retrieved from DB")
// 		} else {
// 			poutpoint.TxId = v
// 		}

// 		pinput.Outpoint = &poutpoint
// 		pinputs = append(pinputs, &pinput)
// 	}

// 	poutputs := make([]*pb.Output, 0, len(tx.Outputs))

// 	for _, output := range tx.Outputs {
// 		var poutput pb.Output

// 		poutput.Value = output.Value
// 		if v, err := hex.DecodeString(output.Pubkey); err != nil {
// 			return nil, errors.New("Output.pubkey is not hex decodable for the object retrieved from DB")
// 		} else {
// 			poutput.Pubkey = v
// 		}
// 		poutputs = append(poutputs, &poutput)
// 	}

// 	pbtx.Type = tx.Type
// 	pbtx.Outputs = poutputs
// 	pbtx.Inputs = pinputs
// 	pbtx.Height = tx.Height

// 	return anypb.New(&pbtx)
// }

// func blkToProtoBlk(in []byte) (*anypb.Any, error) {
// 	var pblk pb.Block
// 	var blk Block

// 	if err := json.Unmarshal(in, &blk); err != nil {
// 		return nil, err
// 	}
// 	pblk.Created = blk.Created
// 	pblk.Miner = blk.Miner
// 	pblk.Type = blk.Type
// 	if v, err := hex.DecodeString(blk.Nonce); err != nil {
// 		return nil, errors.New("Nonce is not hex decodable for the object retrieved from DB")
// 	} else {
// 		pblk.Nonce = v
// 	}

// 	pblk.Note = blk.Note
// 	if v, err := hex.DecodeString(blk.PrevId); err != nil {
// 		return nil, errors.New("PrevId is not hex decodable for the object retrieved from DB")
// 	} else {
// 		pblk.PrevId = v
// 	}

// 	if v, err := hex.DecodeString(blk.T); err != nil {
// 		return nil, errors.New("T is not hex decodable for the object retrieved from DB")
// 	} else {
// 		pblk.T = v
// 	}

// 	var txids [][]byte = make([][]byte, 0)
// 	for _, tx := range blk.TxIds {
// 		if btx, err := hex.DecodeString(tx); err != nil {
// 			return nil, errors.New("Tx Id is not hex decodable for the object retrieved from DB")
// 		} else {
// 			txids = append(txids, btx)
// 		}
// 	}
// 	pblk.TxIds = txids

// 	return anypb.New(&pblk)
// }

// func protoTxToTx(ptx *pb.Tx) ([]byte, error) {
// 	var tx Tx

// 	outputs := make([]*Output, 0)

// 	for _, poutput := range ptx.Outputs {
// 		var output Output
// 		output.Value = poutput.Value
// 		output.Pubkey = hex.EncodeToString(poutput.Pubkey)
// 		outputs = append(outputs, &output)
// 	}
// 	tx.Outputs = outputs

// 	inputs := make([]*Input, 0)

// 	for _, pinput := range ptx.Inputs {
// 		var input Input
// 		input.Sig = pinput.Sig

// 		var outpoint OutPoint
// 		outpoint.Index = pinput.Outpoint.Index
// 		outpoint.TxId = hex.EncodeToString(pinput.Outpoint.TxId)
// 		input.Outpoint = &outpoint
// 		inputs = append(inputs, &input)
// 	}
// 	tx.Inputs = inputs
// 	tx.Type = ptx.Type
// 	tx.Height = ptx.Height

// 	b, err := json.Marshal(tx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return jc.Transform(b)
// }

// func protoBlockToBlock(pblk *pb.Block) ([]byte, error) {
// 	var blk Block
// 	txids := make([]string, 0, len(pblk.TxIds))

// 	//TxIds
// 	for _, tx := range pblk.TxIds {
// 		txStr := hex.EncodeToString(tx)
// 		txids = append(txids, txStr)
// 	}

// 	blk.TxIds = txids
// 	blk.Nonce = hex.EncodeToString(pblk.Nonce)
// 	blk.PrevId = hex.EncodeToString(pblk.PrevId)
// 	blk.Created = pblk.Created
// 	blk.T = hex.EncodeToString(pblk.T)
// 	blk.Miner = pblk.Miner
// 	blk.Note = pblk.Note
// 	blk.Type = pblk.Type

// 	b, err := json.Marshal(blk)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return jc.Transform(b)
// }
