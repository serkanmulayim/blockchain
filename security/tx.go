package security

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"serkanmulayim/blockchain/storage"

	jc "serkanmulayim/blockchain/imports/jsoncanonicalizer"
	pb "serkanmulayim/blockchain/p2p/proto"

	log "github.com/sirupsen/logrus"
)

func HashTxForSignature(tx *pb.Tx) ([]byte, error) {

	signatures := make([][]byte, len(tx.Inputs))
	//set signatures to nil for canonical hash
	for ind, input := range tx.Inputs {
		signatures[ind] = input.Sig
		input.Sig = nil
	}
	//get hash of inputs
	txJson, err := json.Marshal(tx)
	if err != nil { //not expected
		return nil, err
	}
	txCanonicalized, err := jc.Transform(txJson)
	if err != nil { // not expected
		return nil, err
	}
	//reset signatures if there is any
	for ind, input := range tx.Inputs {
		input.Sig = signatures[ind]
	}
	out := sha256.Sum256(txCanonicalized)
	return out[:], nil
}

func ValidateTx(txB []byte) (bool, error) {

	var tx pb.Tx
	err := json.Unmarshal(txB, &tx)

	if err != nil {
		return false, err
	}
	var allInputs, allOutputs int64

	txHash, err := HashTxForSignature(&tx)
	if err != nil {
		return false, err
	}

	//validate outputs
	for _, output := range tx.Outputs {

		if len(output.Pubkey) != ed25519.PublicKeySize {
			return false, errors.New("Output public key is not in good format")
		}
		if output.Value < 0 {
			log.Debugf("Output with pub: %v has a negative value of %d", output.Pubkey, output.Value)
			return false, nil
		}

		allOutputs += output.Value
	}

	//validate inputs
	for _, input := range tx.Inputs {
		//is it in the object store?
		tx_in_byte, err := storage.GetObject(input.Outpoint.TxId, storage.ObjectTypeTx)
		if err != nil || tx_in_byte == nil {
			return false, err
		}

		var tx_in pb.Tx
		err = json.Unmarshal(tx_in_byte, &tx_in)
		//outpoints are associated with outputs
		if input.Outpoint.Index >= int64(len(tx_in.Outputs)) {
			log.Infof("Outpoint index is larger than outputs")
			return false, nil
		}

		//signature: Sign(hash(allTx_withoutSignatures))
		if err != nil || !Verify(tx_in.Outputs[input.Outpoint.Index].Pubkey, txHash, input.Sig) {
			return false, err
		}

		allInputs += tx_in.Outputs[input.Outpoint.Index].Value
	}

	return (allOutputs <= allInputs), nil
}

//non-coinbase txs
// does not do value checking
func SignTx(unsignedTx *pb.Tx, privKey []byte) error {

	hashTx, err := HashTxForSignature(unsignedTx)
	if err != nil {
		return err
	}

	for _, input := range unsignedTx.Inputs {
		input.Sig = Sign(privKey, hashTx)
	}
	return nil
}
