package day_13

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
	sum := 0
	for _, m := range getMaps(input) {
		sum += calculateReflections(m, -1)
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, m := range getMaps(input) {
		sum += calculateReflectionsWithSmudge(m)
	}
	return strconv.Itoa(sum)
}

// getMaps reads the input and returns the separate maps as slice of string slices
func getMaps(input []string) [][]string {
	var r [][]string
	var m []string
	for _, row := range input {
		if len(row) == 0 {
			r = append(r, m)
			m = []string{}
		} else {
			m = append(m, row)
		}
	}
	return append(r, m)
}

// calculateReflections counts the amount of rows (above) and columns (left) reflected on the input map
func calculateReflections(input []string, exclude int) int {
	r := calculateRowReflections(input, exclude/100)
	if r > 0 {
		return 100 * r
	}
	return calculateColumnReflections(input, exclude)
}

// calculateReflectionsWithSmudge counts the amount of rows (above) and columns (left) reflected on the input map
// the only difference if the smudge guessing which tries to flip a single node and recalculate the mirror axis
func calculateReflectionsWithSmudge(input []string) int {
	oldRef := calculateReflections(input, -1)
	for y, row := range input {
		for x, c := range row {
			if c == '#' {
				c = '.'
			} else {
				c = '#'
			}
			newRow := input[y][:x] + string(c) + input[y][x+1:]
			clone := slices.Clone(input)
			clone[y] = newRow
			ref := calculateReflections(clone, oldRef)
			if ref > 0 && ref != oldRef {
				return ref
			}
		}
	}
	panic("no smudge found")
}

// calculateColumnReflections calculates the number of columns reflected on the left side of the mirror
func calculateColumnReflections(input []string, exclude int) int {
	possibleAxis := map[int]int{}
	for _, row := range input {
		stringAxis(row, possibleAxis)
	}
	return getAxis(possibleAxis, len(input), exclude)
}

// calculateRowReflections calculates the number of rows reflected above the mirror
func calculateRowReflections(input []string, exclude int) int {
	possibleAxis := map[int]int{}
	for x := range input[0] {
		column := getColumn(x, input)
		stringAxis(column, possibleAxis)
	}
	return getAxis(possibleAxis, len(input[0]), exclude)
}

// getColumn retrieves the selected column of the input
func getColumn(x int, input []string) string {
	s := ""
	for y := range input {
		s += string(input[y][x])
	}
	return s
}

// stringAxis calculates the overall number of mirroring possibilities for a given string and loads the result into the given map
func stringAxis(s string, possibleAxis map[int]int) {
	for x := 1; x < len(s); x++ {
		left := s[:x]
		right := s[x:]
		ok := true
		for d := 0; d < min(len(left), len(right)) && x-d >= 0; d++ {
			if left[x-d-1] != right[d] {
				ok = false
			}
		}
		if ok {
			possibleAxis[x]++
		}
	}
}

// getAxis finds the common axis within the possible mirrors
func getAxis(possibleAxis map[int]int, length int, exclude int) int {
	for mirror, matches := range possibleAxis {
		if matches == length && mirror != exclude {
			return mirror
		}
	}
	return 0
}
