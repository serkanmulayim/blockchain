package security

import (
	"crypto/ed25519"
)

func Sign(privateKey ed25519.PrivateKey, data []byte) []byte {
	return ed25519.Sign(privateKey, data)

}

func Verify(publicKey []byte, data []byte, sig []byte) bool {
	return ed25519.Verify(publicKey, data, sig)
}

func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}
