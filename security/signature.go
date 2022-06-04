package security

import (
	"crypto/ed25519"
	"crypto/sha256"
)

func Sign(privateKey ed25519.PrivateKey, data []byte) []byte {
	d := sha256.Sum256(data)
	return ed25519.Sign(privateKey, d[:])

}

func Verify(publicKey []byte, data []byte, sig []byte) bool {
	d := sha256.Sum256(data)
	return ed25519.Verify(publicKey, d[:], sig)
}

func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}
