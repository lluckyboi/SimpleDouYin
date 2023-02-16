package tool

import (
	"SimpleDouYin/app/common/key"
	"github.com/cloudwego/thriftgo/pkg/test"
	"testing"
)

func TestHash_Mode(t *testing.T) {
	test.Assert(t, Hash_Mode("testKey", key.RedisHashMod) != "")
}
