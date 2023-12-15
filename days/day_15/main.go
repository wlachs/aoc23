package day_15

import (
	"fmt"
	"strconv"
	"strings"
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
	for _, s := range strings.Split(input[0], ",") {
		sum += int(hash(s))
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// hash calculates the hash of the given string
// The hash is a numerical value from 0 to 255
func hash(s string) uint8 {
	h := uint8(0)
	for _, c := range s {
		h = 17 * (h + uint8(c))
	}
	return h
}
