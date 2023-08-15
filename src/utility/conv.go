package utility

import (
	"strconv"
	"strings"
)

// StringToFloat64 retrieves a float64 value from a given input string
func StringToFloat64(text string) (float64, error) {
	trimmed := strings.Trim(text, " ")
	if len(trimmed) == 0 {
		return 0, nil
	}

	result, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, err
	} else {
		return result, nil
	}
}

// StringToUint64 retrieves a uint64 value from a given input string
func StringToUint64(text string) (uint64, error) {
	trimmed := strings.Trim(text, " ")
	if len(trimmed) == 0 {
		return 0, nil
	}

	result, err := strconv.ParseUint(trimmed, 10, 64)
	if err != nil {
		return 0, err
	} else {
		return result, nil
	}
}

// PascalCase converts a given input string to pascal case
func PascalCase(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
}

// Transform applies the transform function `t` to all elements of `list`
func Transform(list []string, t func(string) string) []string {
	transformed := make([]string, len(list))
	for i, s := range list {
		transformed[i] = t(s)
	}
	return transformed
}
