package tool

import "testing"

func TestLengthCheck(t *testing.T) {
	got := LengthCheck("你好你好你好你好你好11")
	want := true
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = LengthCheck("11")
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = LengthCheck("你好你好你好你好你好111")
	want = false
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = LengthCheck("1")
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
func TestCommentLengthCheck(t *testing.T) {
	got := CommentLengthCheck("你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好111")
	want := true
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = CommentLengthCheck("1")
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = CommentLengthCheck("你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好1111")
	want = false
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	got = CommentLengthCheck("")
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
