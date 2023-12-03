package day_03

import (
	"fmt"
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
	return strconv.Itoa(calculateSumOfParts(input))
}

// calculateSumOfParts calculates the sum of part numbers which are adjacent to a symbol in the input
func calculateSumOfParts(input []string) int {
	sum := 0
	for y, row := range input {
		symbolInRange := false
		buffer := ""
		for x := range row {
			if isNumber(input[y][x]) {
				c := string(input[y][x])
				buffer = buffer + c
				if isSymbolInRange(input, x, y) {
					symbolInRange = true
				}
			} else if symbolInRange {
				symbolInRange = false
				intBuffer, _ := strconv.Atoi(buffer)
				sum += intBuffer
				buffer = ""
			} else {
				buffer = ""
			}
		}
		if symbolInRange {
			intBuffer, _ := strconv.Atoi(buffer)
			sum += intBuffer
			buffer = ""
		}
	}
	return sum
}

// isNumber checks if the given character can be converted to an integer or not
func isNumber(c uint8) bool {
	return c >= 48 && c <= 57
}

// isSymbolInRange checks whether there is a symbol directly before, after, above, below or diagonally next to the current position
func isSymbolInRange(input []string, x0 int, y0 int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if isSymbol(input, x0+x, y0+y) {
				return true
			}
		}
	}
	return false
}

// isSymbol checks if the character at the given position of the input is a symbol or not
func isSymbol(input []string, x int, y int) bool {
	if y < 0 || y >= len(input) {
		return false
	}
	if x < 0 || x >= len(input[0]) {
		return false
	}
	return !isNumber(input[y][x]) && input[y][x] != 46
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
