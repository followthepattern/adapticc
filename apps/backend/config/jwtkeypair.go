package config

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"log/slog"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type JwtKeyPair struct {
	Private crypto.PrivateKey
	Public  crypto.PublicKey
}

func GetKeys(logger *slog.Logger, cfg Server) (JwtKeyPair, error) {
	if len(cfg.Ed25519PrivateKey) < 1 || len(cfg.Ed25519PublicKey) < 1 {
		logger.Info("generating Ed25519 key pairs")
		return GenerateKeys()
	}
	logger.Info("reading Ed25519 key pairs from %s - %s", cfg.Ed25519PrivateKey, cfg.Ed25519PublicKey)
	return ReadKeys(cfg)
}

func GenerateKeys() (JwtKeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return JwtKeyPair{}, err
	}

	return JwtKeyPair{
		Public:  publicKey,
		Private: privateKey,
	}, nil
}

func ReadKeys(cfg Server) (res JwtKeyPair, err error) {
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
