package day_09

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
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
	for _, i := range readInput(input) {
		sum += predictNextValue(i)
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, i := range readInput(input) {
		sum += predictPreviousValue(i)
	}
	return strconv.Itoa(sum)
}

// readInput reads the input and returns it as a 2D int slice
func readInput(input []string) [][]int {
	var l [][]int
	for _, s := range input {
		l = append(l, utils.ToIntSlice(strings.Split(s, " ")))
	}
	return l
}

// predictNextValue predicts the next value of the input based on differences between subsequent values
func predictNextValue(input []int) int {
	if allZero(input) {
		return 0
	}
	var diffs []int
	for i := 0; i < len(input)-1; i++ {
		diffs = append(diffs, input[i+1]-input[i])
	}
	return input[len(input)-1] + predictNextValue(diffs)
}

// predictPreviousValue predicts the previous value of the input based on differences between subsequent values
func predictPreviousValue(input []int) int {
	if allZero(input) {
		return 0
	}
	var diffs []int
	for i := 0; i < len(input)-1; i++ {
		diffs = append(diffs, input[i+1]-input[i])
	}
	return input[0] - predictPreviousValue(diffs)
}

// allZero iterates over the given int slice and checks whether every element of it equals zero
func allZero(diffs []int) bool {
	for _, diff := range diffs {
		if diff != 0 {
			return false
		}
	}
	return true
}
