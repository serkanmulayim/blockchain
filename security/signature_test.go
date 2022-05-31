package security

import "testing"

func TestSignature(t *testing.T) {
	pub, priv, err := GenerateKey()
	if err != nil {
		t.Error("Signing failed")
	}
	data := make([]byte, 100)
	sig := Sign(priv, data[:])

	if !Verify(pub, data[:], sig) {
		t.Error("Verification for correct value failed with false")
	}

	sig[10] = 1

	if Verify(pub, data[:], sig) {
		t.Error("Verification for incorrect signature failed with true")
	}
}
