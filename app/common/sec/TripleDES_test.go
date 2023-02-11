package sec

import (
	"testing"
)

func TestTripleDes(t *testing.T) {
	encrypt, _ := TripleDesEncrypt("secret", "DESKeyDESKeyDESKeyDESKey", "DESIvDES")
	got, _ := TripleDesDecrypt(encrypt, "DESKeyDESKeyDESKeyDESKey", "DESIvDES")
	want := "secret"
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
