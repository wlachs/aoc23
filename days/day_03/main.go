package day_03

import (
	"fmt"
	"slices"
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

// isNothing checks if the given character is "."
func isNothing(c uint8) bool {
	return c == 46
}

// isNumber checks if the given character can be converted to an integer or not
func isNumber(c uint8) bool {
	return c >= 48 && c <= 57
}

// isGear checks if the given character is "*"
func isGear(c uint8) bool {
	return c == 42
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
	return !isNumber(input[y][x]) && !isNothing(input[y][x])
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return strconv.Itoa(calculateSumOfGearParts(input))
}

// calculateSumOfGearParts calculates the sum of products of exactly two part numbers which can be found next to an *
func calculateSumOfGearParts(input []string) int {
	sum := 0
	for y, row := range input {
		for x := range row {
			if isGear(input[y][x]) {
				n := getNeighbouringNumbers(input, x, y)
				if len(n) == 2 {
					sum += n[0] * n[1]
				}
			}
		}
	}
	return sum
}

// getNeighbouringNumbers gets all numbers which have a part neighbouring the current coordinates
func getNeighbouringNumbers(input []string, x0 int, y0 int) []int {
	var n []int
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			i, ok := getNumberIncludingCoordinates(input, x0+x, y0+y)
			if ok && !slices.Contains(n, i) {
				n = append(n, i)
			}
		}
	}
	return n
}

// getNumberIncludingCoordinates checks if the selected coordinates are part of a number and if yes then retrieves it
func getNumberIncludingCoordinates(input []string, x int, y int) (int, bool) {
	if y < 0 || y >= len(input) {
		return 0, false
	}
	if x < 0 || x >= len(input[0]) {
		return 0, false
	}
	if isNumber(input[y][x]) {
		buffer := getNumbersLeft(input[y], x-1) + getNumbersRight(input[y], x)
		i, _ := strconv.Atoi(buffer)
		return i, true
	}
	return 0, false
}

// getNumbersLeft recursively gets a number as string crawling to the left
func getNumbersLeft(s string, x int) string {
	if x < 0 || !isNumber(s[x]) {
		return ""
	}
	return getNumbersLeft(s, x-1) + string(s[x])
}

// getNumbersRight recursively gets a number as string crawling to the right
func getNumbersRight(s string, x int) string {
	if x >= len(s) || !isNumber(s[x]) {
		return ""
	}
	return string(s[x]) + getNumbersRight(s, x+1)
}
