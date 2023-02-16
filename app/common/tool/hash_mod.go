package tool

import (
	"fmt"
	hash2 "github.com/zeromicro/go-zero/core/hash"
)

// Hash_Mode hash取余
func Hash_Mode(key string, mod uint64) string {
	return fmt.Sprintf("_%x", hash2.Hash([]byte(key))%mod)
}
