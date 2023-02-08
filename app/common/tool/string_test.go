package tool

import (
	"testing"
)

func TestRedisStrBuilder(t *testing.T) {
	got := RedisStrBuilder("12.Hh,你")
	want := "12.Hh,你: "
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func TestFiledStringBuild(t *testing.T) {
	slice := []int64{1, 2, 3}
	got := FiledStringBuild("12.Hh,你", slice)
	want := "FIELD(12.Hh,你,1,2,3)"
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func TestAcTypeStringToBool(t *testing.T) {
	got, err := AcTypeStringToBool("1")
	if got != true || err != nil {
		t.Errorf("want:false,nil got:%v,%v", got, err)
	}
	got, err = AcTypeStringToBool("2")
	if got != false || err != nil {
		t.Errorf("want:false,nil got:%v,%v", got, err)
	}
	got, err = AcTypeStringToBool("k")
	if got != true || err == nil {
		t.Errorf("want:true,unknown ActionType got:%v,%v", got, err)
	}
}
