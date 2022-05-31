package security

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"serkanmulayim/blockchain/storage"

	pb "serkanmulayim/blockchain/p2p/proto"

	log "github.com/sirupsen/logrus"
)

func ValidateTx(txB []byte) (bool, error) {

	var tx pb.Tx
	err := json.Unmarshal(txB, &tx)

	if err != nil {
		return false, err
	}

	//validate inputs
	var allInputs, allOutputs int64
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

		b, err := json.Marshal(input.Outpoint)

		if err != nil || !Verify(tx_in.Outputs[input.Outpoint.Index].Pubkey, b, input.Sig) {
			return false, nil
		}

		allInputs += tx_in.Outputs[input.Outpoint.Index].Value
	}

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
	return (allOutputs <= allInputs), nil
}
