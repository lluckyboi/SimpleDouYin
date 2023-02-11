package snowFlake

import (
	"github.com/cloudwego/thriftgo/pkg/test"
	"reflect"
	"testing"
	"time"
)

func TestNextVal(t *testing.T) {
	sf, _ := NewSnowflake(1, 1)
	Id1 := sf.NextVal()
	time.Sleep(1)
	Id2 := sf.NextVal()
	test.Assert(t, Id1 < Id2)
}

func TestGetDeviceID(t *testing.T) {
	sf, _ := NewSnowflake(1, 1)
	d, w := GetDeviceID(sf.NextVal())
	test.Assert(t, d == 1 && w == 1)
}

func TestNewSnowflake(t *testing.T) {
	sf, _ := NewSnowflake(1, 1)
	test.Assert(t, reflect.TypeOf(sf).String() == "*snowFlake.Snowflake")
}
