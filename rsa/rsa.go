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

func ReadPublicKeyFromFile(filePath string) (interface{}, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return x509.ParsePKIXPublicKey(block.Bytes)
}

func ReadPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func Encrypt(plainText []byte, publicKey interface{}) ([]byte, error) {
	v, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key")
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, v, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func EncryptWithFile(plainText []byte, publicKeyFilePath string) ([]byte, error) {
	publicKey, err := ReadPublicKeyFromFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}
	return Encrypt(plainText, publicKey)
}

func Decrypt(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}

func DecryptWithFile(cipherText []byte, privateKeyPath string) ([]byte, error) {
	privateKey, err := ReadPrivateKeyFromFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return Decrypt(cipherText, privateKey)
}
