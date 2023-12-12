package day_12

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"slices"
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
	for _, s := range input {
		row := strings.Split(s, " ")
		d := arrange(row[0], utils.ToIntSlice(strings.Split(row[1], ",")), -1)
		sum += d
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// arrange counts how many different ways can the input be arranged to satisfy the group conditions
func arrange(row string, groups []int, currentGroup int) int {
	if len(row) == 0 && len(groups) == 0 && currentGroup <= 0 {
		return 1
	} else if len(row) == 0 {
		return 0
	}
	switch row[0] {
	case '#':
		if currentGroup == 0 || (currentGroup == -1 && len(groups) == 0) {
			return 0
		} else if currentGroup == -1 {
			currentGroup = groups[0]
			groups = groups[1:]
		}
		return arrange(row[1:], slices.Clone(groups), currentGroup-1)
	case '.':
		if currentGroup <= 0 {
			return arrange(row[1:], slices.Clone(groups), -1)
		}
	case '?':
		return arrange("#"+row[1:], slices.Clone(groups), currentGroup) + arrange("."+row[1:], slices.Clone(groups), currentGroup)
	}
	return 0
}
