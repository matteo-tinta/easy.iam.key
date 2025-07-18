package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"math/big"
)

func base64URLDecode(input string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(input)
}

func main() {
	// Define command-line flags
	nStr := flag.String("n", "", "base64url-encoded modulus (n)")
	eStr := flag.String("e", "", "base64url-encoded exponent (e)")

	flag.Parse()

	if *nStr == "" || *eStr == "" {
		log.Fatalf("Usage: go run main.go -n <modulus> -e <exponent>")
	}

	// Decode n and e
	nBytes, err := base64URLDecode(*nStr)
	if err != nil {
		log.Fatalf("failed to decode n: %v", err)
	}

	eBytes, err := base64URLDecode(*eStr)
	if err != nil {
		log.Fatalf("failed to decode e: %v", err)
	}

	// Convert eBytes to int
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}

	// Convert nBytes to big.Int
	n := new(big.Int).SetBytes(nBytes)

	// Reconstruct RSA public key
	pubKey := rsa.PublicKey{
		N: n,
		E: e,
	}

	// Marshal to ASN.1 DER
	pubASN1, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		log.Fatalf("failed to marshal public key: %v", err)
	}

	// Convert to PEM
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	fmt.Println(string(pubPEM))
}
