package main

import "strconv"

// FormatWithComma adds commas to a numeric string for thousands separators.
func FormatWithComma(s string) string {
	result := ""
	for i, c := range s {
		if (len(s)-i)%3 == 0 && i != 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}

// FormatRank returns a string representing the rank (1-based index) for display.
func FormatRank(rank int) string {
	return FormatWithComma(strconv.Itoa(rank))
}
