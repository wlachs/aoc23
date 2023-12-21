package day_21

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
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
	m := utils.ParseInputToMap(input)
	s := findStart(m)
	return strconv.Itoa(countFields(m, s))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// findStart finds the 'S' node's coordinates on the input map
func findStart(m map[types.Vec2]int32) types.Vec2 {
	for coords, field := range m {
		if field == 'S' {
			return coords
		}
	}
	panic("node not found")
}

// countFields counts the number of reachable fields in 64 steps starting from the given coordinates
func countFields(m map[types.Vec2]int32, s types.Vec2) int {
	acc := []types.Vec2{s}
	for n := 0; n < 64; n++ {
		var nextAcc []types.Vec2
		for _, vec2 := range acc {
			ns := neighbours(m, vec2)
			nextAcc = addIfNotAlready(nextAcc, ns)
		}
		acc = nextAcc
	}
	return len(acc)
}

// neighbours calculates the neighbouring vectors of the given one and checks whether the input map has a navigable field on them
func neighbours(m map[types.Vec2]int32, vec2 types.Vec2) []types.Vec2 {
	res := make([]types.Vec2, 0, 4)
	n := vec2.Around()
	for _, v := range n {
		if m[v] == '.' || m[v] == 'S' {
			res = append(res, v)
		}
	}
	return res
}

// addIfNotAlready adds neighbouring elements to the slice if there are not already contained
func addIfNotAlready(acc []types.Vec2, neighbours []types.Vec2) []types.Vec2 {
	for _, neighbour := range neighbours {
		if !slices.Contains(acc, neighbour) {
			acc = append(acc, neighbour)
		}
	}
	return acc
}
