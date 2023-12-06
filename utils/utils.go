package utils

import "strconv"

// ToIntSlice converts a string slice to an int slice
func ToIntSlice(numbers []string) []int {
	var s []int
	for _, number := range numbers {
		s = append(s, Atoi(number))
	}
	return s
}

// Atoi converts a string to an int without the error return value.
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
