package day_12

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
	return strconv.Itoa(calculateSum(input, 1))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return strconv.Itoa(calculateSum(input, 5))
}

// calculateSum calculates the overall sum of different arrangement possibilities across all input rows
func calculateSum(input []string, foldingFactor int) int {
	sum := 0
	cache := map[string]int{}
	for _, s := range input {
		records, conditions := processInput(s, foldingFactor)
		sum += arrange(records, conditions, -1, cache)
	}
	return sum
}

// processInput parses a single input row with the given folding factor.
// The folding factor defines how many times the input string and the spring condition should be repeated
func processInput(input string, foldingFactor int) (string, []int) {
	row := strings.Split(input, " ")
	r := row[0]
	c := utils.ToIntSlice(strings.Split(row[1], ","))
	records := r
	conditions := c

	for i := 1; i < foldingFactor; i++ {
		records += "?" + r
		conditions = append(conditions, c...)
	}

	return records, conditions
}

// arrange counts how many different ways can the input be arranged to satisfy the group conditions
func arrange(row string, groups []int, currentGroup int, cache map[string]int) int {
	m, ok := cache[genKey(row, groups, currentGroup)]
	if ok {
		return m
	}
	if len(row) == 0 && len(groups) == 0 && currentGroup <= 0 {
		return 1
	} else if len(row) == 0 {
		return 0
	}
	d := 0
	switch row[0] {
	case '#':
		if currentGroup == 0 || (currentGroup == -1 && len(groups) == 0) {
			return 0
		} else if currentGroup == -1 {
			currentGroup = groups[0]
			groups = groups[1:]
		}
		d = arrange(row[1:], groups, currentGroup-1, cache)
	case '.':
		if currentGroup <= 0 {
			d = arrange(row[1:], groups, -1, cache)
		}
	case '?':
		d = arrange("#"+row[1:], groups, currentGroup, cache) + arrange("."+row[1:], groups, currentGroup, cache)
	}
	cache[genKey(row, groups, currentGroup)] = d
	return d
}

// genKey generates a cache key to avoid redundant calculations
func genKey(row string, groups []int, currentGroup int) string {
	return row + "->" + strings.Join(utils.ToStringSlice(groups), ",") + "->" + strconv.Itoa(currentGroup)
}
