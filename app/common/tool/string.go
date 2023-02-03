package tool

import (
	"errors"
	"strconv"
	"strings"
)

func RedisStrBuilder(key string) string {
	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString(": ")
	return builder.String()
}

func FiledStringBuild(column string, IDs []int64) string {
	var strb strings.Builder
	strb.WriteString("FIELD(")
	strb.WriteString(column)
	for _, val := range IDs {
		strb.WriteString(",")
		strb.WriteString(strconv.FormatInt(val, 10))
	}
	strb.WriteString(")")
	return strb.String()
}

func AcTypeStringToBool(act string) (bool, error) {
	switch act {
	case "1":
		return true, nil
	case "2":
		return false, nil
	default:
		return true, errors.New("unknown ActionType")
	}
}
