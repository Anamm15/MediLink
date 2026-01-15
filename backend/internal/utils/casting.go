package utils

import (
	"strconv"
	"strings"
)

func StringToInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	return strconv.Atoi(s)
}
