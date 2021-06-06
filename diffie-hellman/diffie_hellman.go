package diffiehellman

import (
	"crypto/rand"
	"log"
	"math/big"
)

// PrivateKey creates a random number in the range (1, mod)
// mod must be a prime number.
func PrivateKey(mod *big.Int) *big.Int {
	one := new(big.Int).SetInt64(1)

	for {
		// rand.Int generates a number in the range [0,modulus)
		priv, err := rand.Int(rand.Reader, mod)
		if err != nil {
			log.Fatal(err)
		}

		// make sure our number is > 1
		if priv.Cmp(one) > 0 {
			return priv
		}
	}
}

// PublicKey creates the public key for the given
// private key, modulus, and base. base should be a primitive root modulo mod
func PublicKey(priv, mod *big.Int, gen int64) *big.Int {
	g := new(big.Int).SetInt64(gen)
	return new(big.Int).Exp(g, priv, mod)
}

// NewPair creates a new private/public key pair
func NewPair(mod *big.Int, gen int64) (priv, pub *big.Int) {
	priv = PrivateKey(mod)
	pub = PublicKey(priv, mod, gen)
	return
}

// SecretKey creates a new shared key.  This function should be given
// Alice's private key, Bob's public key, and the agreed-upon modulus.
// It will return a secret key that only Alice and Bob will be able to know.
func SecretKey(priv, pub, mod *big.Int) *big.Int {
	return new(big.Int).Exp(pub, priv, mod)
}
