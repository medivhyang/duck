package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

var ErrInvalidPublicKey = errors.New("invalid rsa public key")

func GenerateKeyPairFiles(bits int, publicKeyFilePath, privateKeyFilePath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyFile, err := os.Create(privateKeyFilePath)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()
	privateKeyBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	if err = pem.Encode(privateKeyFile, &privateKeyBlock); err != nil {
		return err
	}

	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	publicKeyFile, err := os.Create(publicKeyFilePath)
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()
	publicKeyBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	if err := pem.Encode(publicKeyFile, &publicKeyBlock); err != nil {
		return err
	}

	return nil
}

func ParsePublicKeyFromPem(b []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(b)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			key = cert.PublicKey
		} else {
			return nil, err
		}
		return nil, err
	}
	v, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, ErrInvalidPublicKey
	}
	return v, nil
}

func ParsePrivateKeyFromPem(b []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func Encrypt(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func Decrypt(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}

func EncryptWithBytes(plainText []byte, publicKeyBytes []byte) ([]byte, error) {
	publicKey, err := ParsePublicKeyFromPem(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func DecryptWithBytes(cipherText []byte, privateKeyBytes []byte) ([]byte, error) {
	privateKey, err := ParsePrivateKeyFromPem(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}

func EncryptWithFile(plainText []byte, publicKeyFilePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}
	return EncryptWithBytes(plainText, b)
}

func DecryptWithFile(cipherText []byte, privateKeyFilePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return nil, err
	}
	return DecryptWithBytes(cipherText, b)
}
