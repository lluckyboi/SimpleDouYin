package tool

import (
	"testing"
)

func TestFilter(t *testing.T) {
	st := NewSensitiveTrie('*').Init([]string{"傻"})
	_, res := st.Match("好傻")
	if res != "好*" {
		t.Errorf("fail")
	}
}
