package tool

import (
	"strings"
	"testing"
)

func TestLengthCheck(t *testing.T) {
	type args struct {
		str string
	}

	var bd strings.Builder
	for i := 0; i < 32; i++ {
		bd.WriteString("0")
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"str with len 0", args{str: ""}, false},
		{"str with len 1", args{str: "s"}, false},
		{"str with len 2", args{str: "st"}, true},
		{"str with len 3", args{str: "str"}, true},
		{"str with len 32", args{str: bd.String()}, true},
		{"str with len 33", args{str: bd.String() + "0"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthCheck(tt.args.str); got != tt.want {
				t.Errorf("LengthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentLengthCheck(t *testing.T) {
	type args struct {
		str string
	}

	var bd strings.Builder
	for i := 0; i < 255; i++ {
		bd.WriteString("0")
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"str with len 0", args{str: ""}, false},
		{"str with len 1", args{str: "s"}, true},
		{"str with len 255", args{str: bd.String()}, true},
		{"str with len 256", args{str: bd.String() + "0"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommentLengthCheck(tt.args.str); got != tt.want {
				t.Errorf("LengthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
