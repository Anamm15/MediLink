package utils

import (
	"strconv"
	"strings"
)

func StringToInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	return strconv.Atoi(s)
}

func StringToBoolDefault(s string, defaultVal bool) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return defaultVal
	}

	val, err := strconv.ParseBool(s)
	if err != nil {
		return defaultVal
	}
	return val
}
