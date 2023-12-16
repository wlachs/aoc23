package utils

import "strconv"

// Atoi converts a string to an int without the error return value.
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
