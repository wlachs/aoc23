package utils

import "strconv"

// ToIntSlice converts a string slice to an int slice
func ToIntSlice(numbers []string) []int {
	s := make([]int, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, Atoi(number))
	}
	return s
}

// ToStringSlice converts an int slice to a string slice
func ToStringSlice(numbers []int) []string {
	s := make([]string, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, strconv.Itoa(number))
	}
	return s
}
