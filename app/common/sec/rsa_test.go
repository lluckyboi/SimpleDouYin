package sec

import (
	"testing"
)

func TestRSA(t *testing.T) {
	cipherText := RSA_Encrypt([]byte("Zhishixuebao123"), []byte("114514"), false)
	got := RSA_Decrypt([]byte(cipherText), []byte("114514"), false)
	want := "Zhishixuebao123"
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
