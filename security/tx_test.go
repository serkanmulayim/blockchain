package security

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"serkanmulayim/blockchain/storage"
	"testing"

	jc "serkanmulayim/blockchain/imports/jsoncanonicalizer"
	pb "serkanmulayim/blockchain/p2p/proto"

	log "github.com/sirupsen/logrus"
)

const (
	port     = 11
	coinbase = `{"height":0,"outputs":[{
	"pubkey":"8dbcd2401c89c04d6e53c81c90aa0b551cc8fc47c0469217c8f5cfbae1e911f9",
	"value":50000000000}],"type":"tx"}`

	tx1 = `{"inputs":[{"outpoint":{"index":0,
	"txid":"1bb37b637d07100cd26fc063dfd4c39a7931cc88dae3417871219715a5e374af"},
	"sig":"1d0d7d774042607c69a87ac5f1cdf92bf474c25fafcc089fe667602bfefb0494726c519e92266957429ced875256e6915eb8cea2ea66366e739415efc47a6805"}],
	"outputs":[{
	"pubkey":"8dbcd2401c89c04d6e53c81c90aa0b551cc8fc47c0469217c8f5cfbae1e911f9",
	"value":10}],"type":"tx"}`
)

func start() {
	storage.StartLevelDB(port)

}

func teardown() {
	storage.StopLevelDb()

	dir := fmt.Sprintf("./db/%d%v.db", port, storage.PeersPrefix)
	err := os.RemoveAll(dir)
	if err != nil {
		log.Errorf("%v", err)
	}
	dir = fmt.Sprintf("./db/%d%v.db", port, storage.BlockObjectsPrefix)
	err = os.RemoveAll(dir)
	if err != nil {
		log.Errorf("%v", err)
	}
	dir = fmt.Sprintf("./db/%d%v.db", port, storage.TxObjectsPrefix)
	err = os.RemoveAll(dir)
	if err != nil {
		log.Errorf("%v", err)
	}
}

func TestTx1(t *testing.T) {
	start()
	defer teardown()
	var cbTx pb.Tx
	var txTx pb.Tx

	cbPub, cbPriv, err := GenerateKey()
	txPub, _, err := GenerateKey()

	cbTx.Height = 0
	var cbOutput pb.Output
	cbOutput.Pubkey = cbPub
	cbOutput.Value = 5000000000000000 // 5 * 10^15
	cbTx.Outputs = append(cbTx.Outputs, &cbOutput)
	cbTxBytes, err := json.Marshal(cbTx)
	if err != nil {
		t.Error("could not create the Coinbase transaction")
	}

	cbTxCanonical, err := jc.Transform(cbTxBytes)
	if err != nil {
		t.Error("cannot canonicalize CB")
	}

	cbObjId := sha256.Sum256(cbTxCanonical)

	err = storage.PutObject(cbTxBytes, storage.ObjectTypeTx)
	if err != nil {
		t.Error("cannot put Coinbase to DB")
	}

	var outpoint pb.OutPoint
	var input pb.Input
	outpoint.Index = 0
	outpoint.TxId = cbObjId[:]
	outpointBytes, err := json.Marshal(outpoint)
	if err != nil {
		t.Error("cannot marshal outpointBytes")
	}

	outpointCanonical, err := jc.Transform(outpointBytes)
	if err != nil {
		t.Error("Could not canonicalize outpoint")
	}

	input.Sig = Sign(cbPriv, outpointCanonical)
	input.Outpoint = &outpoint

	txTx.Inputs = append(txTx.Inputs, &input)

	var output1 pb.Output
	var output2 pb.Output

	output1.Pubkey = cbPub
	output1.Value = 4000000000000000

	output2.Pubkey = txPub
	output2.Value = 1000000000000000
	txTx.Outputs = append(txTx.Outputs, &output1, &output2)

	//valid
	txBytesValid, err := json.Marshal(txTx)
	if err != nil {
		t.Error("could not marshal Tx")
	}

	out, err := ValidateTx(txBytesValid)
	if !out {
		t.Error("Valid transaction could not be validated")
	}

	//input < output
	output2.Value = 2000000000000000
	txBytesValid, err = json.Marshal(txTx)
	if err != nil {
		t.Error("could not marshal Tx")
	}

	out, err = ValidateTx(txBytesValid)
	if out {
		t.Error("Invalid transaction with inputs<outputs is validated")
	}

	//bad signature
	output2.Value = 1000000000000000
	randomSignature := make([]byte, ed25519.SignatureSize)
	rand.Read(randomSignature)
	txTx.Inputs[0].Sig = randomSignature
	txBytesValid, err = json.Marshal(txTx)
	if err != nil {
		t.Error("could not marshal Tx")
	}

	out, err = ValidateTx(txBytesValid)
	if out {
		t.Error("Invalid transaction with bad input signature is validated")
	}
}
