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

// AsEnumDefinition uppers the text and replaces space with underscores
func AsEnumDefinition(text string) string {
	return strings.ReplaceAll(strings.ToUpper(text), " ", "_")
}

// Transform applies the transform function `t` to all elements of `list`
func Transform(list []string, t func(string) string) []string {
	transformed := make([]string, len(list))
	for i, s := range list {
		transformed[i] = t(s)
	}
	return transformed
}
