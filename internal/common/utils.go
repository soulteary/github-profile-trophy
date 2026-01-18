package common

import (
	"math"
	"strconv"
	"strings"
)

// AbridgeScore formats a score with "pt" suffix, using "k" for thousands
func AbridgeScore(score int) string {
	absScore := math.Abs(float64(score))
	if absScore < 1 {
		return "0pt"
	}
	if absScore > 999 {
		sign := 1
		if score < 0 {
			sign = -1
		}
		return strconv.FormatFloat(float64(sign)*absScore/1000, 'f', 1, 64) + "kpt"
	}
	return strconv.Itoa(score) + "pt"
}

// ParseBoolean returns boolean if value is either "true" or "false" else returns nil
func ParseBoolean(value interface{}) *bool {
	if b, ok := value.(bool); ok {
		return &b
	}

	if str, ok := value.(string); ok {
		lower := strings.ToLower(str)
		if lower == "true" {
			b := true
			return &b
		} else if lower == "false" {
			b := false
			return &b
		}
	}
	return nil
}

// ParseArray parses string to array of strings
func ParseArray(str string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}
