package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
)

type RSAKeyInfo struct {
	PrivateKey string `json:"private_key"` // base64 of PEM
	PublicKey  string `json:"public_key"`  // base64 of PEM
	N          string `json:"n"`           // modulus
	E          string `json:"e"`           // exponent
	Alg        string `json:"alg"`         // algoritm
}

// intToBytes converts an int to a big-endian byte slice
func intToBytes(i int) []byte {
	return new(big.Int).SetInt64(int64(i)).Bytes()
}

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Encode private key to PEM
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	// Encode public key to PEM
	pubASN1, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		panic(err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	// Base64-encode the full PEM text
	privBase64 := base64.StdEncoding.EncodeToString(privPEM)
	pubBase64 := base64.StdEncoding.EncodeToString(pubPEM)

	// Convert N and E to base64url
	nBase64URL := base64.RawURLEncoding.EncodeToString(key.PublicKey.N.Bytes())
	eBase64URL := base64.RawURLEncoding.EncodeToString(intToBytes(key.PublicKey.E))

	keyInfo := RSAKeyInfo{
		PrivateKey: privBase64,
		PublicKey:  pubBase64,
		N:          nBase64URL,
		E:          eBase64URL,
		Alg:        "RS256",
	}

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(keyInfo, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
