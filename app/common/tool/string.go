package tool

import (
	"strings"
)

func RedisStrBuilder(key string) string {
	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString(": ")
	return builder.String()
}
