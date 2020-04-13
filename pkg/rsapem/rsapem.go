package rsapem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const rsaPrivateKey = "RSA PRIVATE KEY"

// RSAPrivateToPem converts rsa private key into a PEM block
func RSAPrivateToPem(privateKey *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  rsaPrivateKey,
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
}

// RSAPrivateFromPen decodes a private key from a PEM file
func RSAPrivateFromPen(data []byte)(*rsa.PrivateKey, error) {
	p, _ := pem.Decode(data)
	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
