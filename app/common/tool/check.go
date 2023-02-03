package tool

func LengthCheck(ss string) bool {
	if len(ss) > 32 || len(ss) < 2 {
		return false
	}
	return true
}

func CommentLengthCheck(ss string) bool {
	if len(ss) > 320 || len(ss) < 1 {
		return false
	}
	return true
}
