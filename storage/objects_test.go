package storage

// import (
// 	"encoding/hex"
// 	"encoding/json"
// 	"reflect"
// 	"testing"

// 	pb "serkanmulayim/blockchain/p2p/proto"

// 	"google.golang.org/protobuf/types/known/anypb"
// )

// func TestBlock(t *testing.T) {

// 	//1st way Block->Any->Json
// 	var pblk pb.Block

// 	pblk.Created = 1000
// 	pblk.Miner = "serkan"
// 	pblk.Nonce = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	pblk.Note = "this is a note"
// 	pblk.PrevId = []byte{0, 1, 2, 3, 4, 5, 6, 7}
// 	pblk.T = []byte{0, 1, 2, 3, 4, 5}
// 	pblk.TxIds = [][]byte{{0, 0, 0, 0}, {1, 1, 1, 1}, {2, 2, 2, 2}}
// 	pblk.Type = ObjectTypeBlock

// 	any, err := anypb.New(&pblk)
// 	if err != nil {
// 		t.Error("Cannot convert pb.Block to any")
// 	}

// 	bytesJson, err := ProtoToJson(any)

// 	var get Block
// 	var expectedBlk Block

// 	expectedBlk.Created = pblk.Created
// 	expectedBlk.Miner = pblk.Miner
// 	expectedBlk.Nonce = hex.EncodeToString(pblk.Nonce)
// 	expectedBlk.Note = pblk.Note
// 	expectedBlk.PrevId = hex.EncodeToString(pblk.PrevId)
// 	expectedBlk.T = hex.EncodeToString(pblk.T)
// 	expectedBlk.TxIds = make([]string, 0, len(pblk.TxIds))
// 	expectedBlk.Type = ObjectTypeBlock

// 	for _, arr := range pblk.TxIds {
// 		expectedBlk.TxIds = append(expectedBlk.TxIds, hex.EncodeToString(arr))
// 	}
// 	pblk.Type = ObjectTypeBlock

// 	err = json.Unmarshal(bytesJson, &get)
// 	if err != nil {
// 		t.Error("Error in unmarshaling json to Block")
// 	}

// 	if !reflect.DeepEqual(get, expectedBlk) {
// 		t.Error("JSON Block Objects are not equal")
// 	}

// 	//backward json->any->Block

// 	any2, _ := JsonToProto(bytesJson)

// 	msg, _ := any2.UnmarshalNew()
// 	pblk2 := msg.(*pb.Block)

// 	if !pbBlockEquals(pblk, *pblk2) {
// 		t.Error("PB.Block objects are not equal")
// 	}
// }

// func TestTx(t *testing.T) {
// 	var ptx pb.Tx

// 	ptx.Type = ObjectTypeTx
// 	ptx.Height = 12

// 	//INPUTS
// 	var pinputs []*pb.Input = make([]*pb.Input, 0)
// 	//pinput1
// 	var pinput1 pb.Input
// 	pinput1.Sig = []byte{0, 1, 2, 3, 4, 5}
// 	pinput1.Outpoint = &pb.OutPoint{Index: 0, TxId: []byte{2, 2, 2, 2, 2, 2}}
// 	pinputs = append(pinputs, &pinput1)

// 	//pinput2
// 	var pinput2 pb.Input
// 	pinput2.Sig = []byte{10, 11, 12, 13, 14, 15}
// 	pinput2.Outpoint = &pb.OutPoint{Index: 1, TxId: []byte{12, 12, 12, 12, 12, 12}}
// 	pinputs = append(pinputs, &pinput2)

// 	ptx.Inputs = pinputs

// 	//OUTPUTS
// 	var poutputs []*pb.Output = make([]*pb.Output, 0)
// 	//poutput1
// 	var poutput1 pb.Output
// 	poutput1.Pubkey = []byte{1, 1, 1, 1, 1, 1}
// 	poutput1.Value = 1

// 	var poutput2 pb.Output
// 	poutput2.Pubkey = []byte{2, 2, 2, 2, 2, 2}
// 	poutput2.Value = 2

// 	poutputs = append(poutputs, &poutput1, &poutput2)
// 	ptx.Outputs = poutputs

// 	//Start test

// 	any, err := anypb.New(&ptx)
// 	if err != nil {
// 		t.Errorf("Failure in anypb.New: %v", err)
// 	}

// 	bytesJson, err := ProtoToJson(any)

// 	var gotTx Tx
// 	json.Unmarshal(bytesJson, &gotTx)

// 	var expectedTx Tx

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

// 	expectedTx.Inputs = inputs

// 	outputs := make([]*Output, 0)

// 	for _, poutput := range ptx.Outputs {
// 		var output Output
// 		output.Pubkey = hex.EncodeToString(poutput.Pubkey)
// 		output.Value = poutput.Value
// 		outputs = append(outputs, &output)
// 	}

// 	expectedTx.Outputs = outputs
// 	expectedTx.Type = ObjectTypeTx
// 	expectedTx.Height = 12

// 	if !reflect.DeepEqual(expectedTx, gotTx) {
// 		t.Error("Received Tx is not the same as expected one")
// 	}

// 	//backward json->any->Block

// 	any2, _ := JsonToProto(bytesJson)

// 	msg, _ := any2.UnmarshalNew()
// 	ptx2 := msg.(*pb.Tx)

// 	if !pbTxEquals(ptx, *ptx2) {
// 		t.Error("PB.Block objects are not equal")
// 	}

// }

// func pbTxEquals(exp pb.Tx, got pb.Tx) bool {
// 	out := (exp.Type == got.Type)
// 	out = out && inputsEquals(exp.Inputs, got.Inputs)
// 	out = out && outputsEquals(exp.Outputs, exp.Outputs)

// 	return out
// }

// func outputsEquals(exp []*pb.Output, got []*pb.Output) bool {
// 	if len(exp) != len(got) {
// 		return false
// 	}

// 	for i, goutput := range got {
// 		if goutput.Value != exp[i].Value {
// 			return false
// 		}
// 		if !reflect.DeepEqual(goutput.Pubkey, exp[i].Pubkey) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func inputsEquals(exp []*pb.Input, got []*pb.Input) bool {
// 	if len(exp) != len(got) {
// 		return false
// 	}
// 	for i, ginput := range got {
// 		if !reflect.DeepEqual(ginput.Sig, exp[i].Sig) {
// 			return false
// 		}
// 		if ginput.Outpoint.Index != exp[i].Outpoint.Index {
// 			return false
// 		}
// 		if !reflect.DeepEqual(ginput.Outpoint.TxId, exp[i].Outpoint.TxId) {
// 			return false
// 		}
// 	}
// 	return true

// }

// func pbBlockEquals(exp pb.Block, get pb.Block) bool {
// 	out := true

// 	out = out && (exp.Created == get.Created)
// 	out = out && (exp.Miner == get.Miner)
// 	out = out && reflect.DeepEqual(exp.Nonce, get.Nonce)
// 	out = out && (exp.Note == get.Note)
// 	out = out && reflect.DeepEqual(exp.PrevId, get.PrevId)
// 	out = out && reflect.DeepEqual(exp.T, get.T)
// 	out = out && (exp.Type == get.Type)
// 	out = out && reflect.DeepEqual(exp.TxIds, get.TxIds)

// 	return out
// }
