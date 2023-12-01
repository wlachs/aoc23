package day_01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mapping = map[string]int{
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	sum := 0
	for _, line := range input {
		sum += firstInt(line)*10 + lastInt(line)
	}
	return strconv.Itoa(sum)
}

// firstInt gets the first integer in a given string
func firstInt(s string) int {
	re := regexp.MustCompile("^[a-z]*\\d")
	match := re.FindString(s)
	return int(match[len(match)-1]) - 48
}

// lastInt gets the last integer in a given string
func lastInt(s string) int {
	re := regexp.MustCompile("\\d[a-z]*$")
	match := re.FindString(s)
	return int(match[0]) - 48
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, line := range input {
		sum += firstIntExtended(line)*10 + lastIntExtended(line)
	}
	return strconv.Itoa(sum)
}

// firstIntExtended gets the first integer in a given string and also matches words
func firstIntExtended(s string) int {
	for key, value := range mapping {
		if strings.HasPrefix(s, key) {
			return value
		}
	}
	return firstIntExtended(s[1:])
}

// lastIntExtended gets the last integer in a given string and also matches words
func lastIntExtended(s string) int {
	for key, value := range mapping {
		if strings.HasSuffix(s, key) {
			return value
		}
	}
	return lastIntExtended(s[0 : len(s)-1])
}
