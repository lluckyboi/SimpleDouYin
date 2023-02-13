package tool

import (
	"reflect"
	"testing"
)

func TestInitLogger(t *testing.T) {
	logger := InitLogger()
	switch reflect.TypeOf(logger).String() {
	case "*zap.Logger":
	default:
		t.Errorf("init err")
	}
}
