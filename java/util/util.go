package util

import (
	"fmt"
	"strconv"
	"strings"
)

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func StrFormat(format string, params ...interface{}) string {
	paramLength := len(params)
	length := len(format)
	paramIndex := 0
	var build strings.Builder
	for i := 0; i < length; i++ {
		char := string(format[i])
		if char == "{" && i != length-1 {
			nextChar := string(format[i+1])
			if nextChar == "}" {
				if paramIndex < paramLength {
					build.WriteString(getStringParam(params[paramIndex]))
				} else {
					build.WriteString("{}")
				}
				paramIndex++
				i++
			} else {
				build.WriteString("{")
			}
		} else {
			build.WriteString(char)
		}
	}
	return build.String()
}
func getStringParam(v interface{}) string {
	switch v.(type) {
	case int:
		return strconv.Itoa(v.(int))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case int32:
		return strconv.FormatInt(int64(v.(int32)), 10)
	case string:
		return v.(string)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	default:
		return fmt.Sprintf("%c", v)
	}
}
