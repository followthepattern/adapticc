package config

import (
	"crypto"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type JwtKeyPair struct {
	Private crypto.PrivateKey
	Public  crypto.PublicKey
}

func ReadKeys(cfg Server) (JwtKeyPair, error) {
	res := JwtKeyPair{}

	key, err := os.ReadFile(cfg.Ed25519PublicKey)
	if err != nil {
		return JwtKeyPair{}, err
	}

	ed25519PublicKey, err := jwt.ParseEdPublicKeyFromPEM(key)
	if err != nil {
		return JwtKeyPair{}, err
	}

	key, err = os.ReadFile(cfg.Ed25519PrivateKey)
	if err != nil {
		return JwtKeyPair{}, err
	}

	ed25519PrivateKey, err := jwt.ParseEdPrivateKeyFromPEM(key)
	if err != nil {
		return res, err
	}

	return JwtKeyPair{
		Public:  ed25519PublicKey,
		Private: ed25519PrivateKey,
	}, nil
}
