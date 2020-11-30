package rsa

import "testing"

func TestGenerateKeyPairFiles(t *testing.T) {
	if err := GenerateKeyPairFiles(2048, "public.pem", "private.pem"); err != nil {
		t.Error(err)
	}
}
