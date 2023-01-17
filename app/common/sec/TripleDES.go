package sec

import (
	"encoding/base64"
	"github.com/wumansgy/goEncrypt/des"
)

// TripleDesEncrypt 三重DES加密
func TripleDesEncrypt(encryptedString string, DESKey string, DESIv string) (string, error) {
	cryptText, err := des.TripleDesEncrypt([]byte(encryptedString), []byte(DESKey), []byte(DESIv))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(cryptText), nil
}

// TripleDesDecrypt 三重DES解密
func TripleDesDecrypt(decryptString string, DESKey string, DESIv string) (string, error) {
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptString)
	if err != nil {
		return "", err
	}
	cryptText, err := des.TripleDesDecrypt(decryptBytes, []byte(DESKey), []byte(DESIv))
	if err != nil {
		return "", err
	}

	return string(cryptText), nil
}
