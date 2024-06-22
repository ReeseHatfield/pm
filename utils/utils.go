package utils

import "strings"

func IsValidYesNoString(s string) bool {
	return strings.HasPrefix(s, "y") || strings.HasPrefix(s, "n")
}
