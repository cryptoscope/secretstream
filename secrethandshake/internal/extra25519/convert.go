// blatent grep of github.com/agl/ed25519/extra25519 until we have a better solution
package extra25519

import (
	"crypto/sha512"

	"filippo.io/edwards25519"
	"golang.org/x/crypto/ed25519"

	"go.cryptoscope.co/secretstream/internal/lo25519"
)

// PrivateKeyToCurve25519 converts an ed25519 private key into a corresponding
// curve25519 private key such that the resulting curve25519 public key will
// equal the result from PublicKeyToCurve25519.
func PrivateKeyToCurve25519(curve25519Private *[32]byte, privateKey ed25519.PrivateKey) {
	h := sha512.New()
	h.Write(privateKey[:32])
	digest := h.Sum(nil)

	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(curve25519Private[:], digest)
}

// PublicKeyToCurve25519 converts an Ed25519 public key into the curve25519
// public key that would be generated from the same private key.
func PublicKeyToCurve25519(curve25519Public *[32]byte, publicKey ed25519.PublicKey) bool {
	var ge [32]byte
	copy(ge[:], publicKey)
	if lo25519.IsEdLowOrder(ge[:]) {
		return false
	}

	a, err := new(edwards25519.Point).SetBytes(ge[:])
	if err != nil {
		return false
	}

	copy(curve25519Public[:], a.BytesMontgomery())
	return true
}
