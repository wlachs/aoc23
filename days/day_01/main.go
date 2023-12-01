package day_01

import (
	"fmt"
	"regexp"
	"strconv"
)

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
	return ""
}
