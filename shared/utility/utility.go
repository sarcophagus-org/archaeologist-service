package utility

import "strings"

func IsHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}