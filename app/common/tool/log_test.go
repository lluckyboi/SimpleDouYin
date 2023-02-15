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

func TestLogConstruct(t *testing.T) {
	_, err := LogConstruct("info", "", 1, "dsa")
	if err != nil {
		t.Errorf("error:%v", err)
	}
}
