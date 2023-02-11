package sec

import (
	"testing"
)

func TestRSA(t *testing.T) {
	cipherText := RSA_Encrypt([]byte("Zhishixuebao123"), []byte("-----BEGIN RSA Public Key-----\nMDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhAMwhE5MYrM+zw8pus8+SOeSSjUx8OB3x\nCpu9mkTOL6QzAgMBAAE=\n-----END RSA Public Key-----"), false)
	got := RSA_Decrypt([]byte(cipherText), []byte("-----BEGIN RSA Private Key-----\nMIGqAgEAAiEAzCETkxisz7PDym6zz5I55JKNTHw4HfEKm72aRM4vpDMCAwEAAQIh\nAI94V/e1ChDZuizXbc3gaor5Caab/EkRY/yi8+ShjdyRAhEA4IePbB162RTTvv3a\nZGYHlQIRAOi9hfzkPkoS5HeWsYtJaqcCEBEG9r9yNODFjZFMWwWGH0kCEDrIy9Ph\nLl51QSF3fWaJ55cCECuLHotkGtTRXOLMexpimV4=\n-----END RSA Private Key-----\n"), false)
	want := ""
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
