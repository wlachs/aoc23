package day_23

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
	for x := 0; x < len(input[0]); x++ {
		pos := types.Vec2{X: x}
		if m[pos] == '.' {
			return strconv.Itoa(findLongestPath(m, []types.Vec2{pos}, types.Vec2{X: len(input[0]), Y: len(input)}, false))
		}
	}
	panic("no starting field found")
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	for x := 0; x < len(input[0]); x++ {
		pos := types.Vec2{X: x}
		if m[pos] == '.' {
			return strconv.Itoa(findLongestPath(m, []types.Vec2{pos}, types.Vec2{X: len(input[0]), Y: len(input)}, true))
		}
	}
	panic("no starting field found")
}

var best = 0

// findLongestPath iterates over the input map and finds the longest hiking path without loops
func findLongestPath(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2, ignoreSlopes bool) int {
	var options []types.Vec2
	for len(options) < 2 {
		if len(options) == 1 {
			path = append(path, options[0])
		}
		if ignoreSlopes {
			options = findNonSlipperyNextOptions(m, path, dim)
		} else {
			options = findNextOptions(m, path, dim)
		}
		if len(options) == 0 {
			if path[len(path)-1].Y == dim.Y-1 {
				return len(path) - 1
			}
			return 0
		}
	}
	longest := 0
	for _, option := range options {
		newPath := slices.Clone(path)
		newPath = append(newPath, option)
		longest = max(longest, findLongestPath(m, newPath, dim, ignoreSlopes))
		if longest > best {
			best = longest
			fmt.Println(best)
		}
	}
	return longest
}

// findNextOptions checks the surrounding fields for the next step of the hike
func findNextOptions(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2) []types.Vec2 {
	options := make([]types.Vec2, 0, 4)
	pos := path[len(path)-1]
	up := pos.Up()
	if up.Y >= 0 && m[up] != 'v' && m[up] != '#' && !slices.Contains(path, up) {
		options = append(options, up)
	}
	down := pos.Down()
	if down.Y < dim.Y && m[down] != '^' && m[down] != '#' && !slices.Contains(path, down) {
		options = append(options, down)
	}
	left := pos.Left()
	if left.X >= 0 && m[left] != '>' && m[left] != '#' && !slices.Contains(path, left) {
		options = append(options, left)
	}
	right := pos.Right()
	if right.X < dim.X && m[right] != '<' && m[right] != '#' && !slices.Contains(path, right) {
		options = append(options, right)
	}
	return options
}

// findNonSlipperyNextOptions checks the surrounding fields for the next step of the hike
// this variant of the method ignores steep slopes
func findNonSlipperyNextOptions(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2) []types.Vec2 {
	options := make([]types.Vec2, 0, 4)
	pos := path[len(path)-1]
	up := pos.Up()
	if up.Y >= 0 && m[up] != '#' && !slices.Contains(path, up) {
		options = append(options, up)
	}
	down := pos.Down()
	if down.Y < dim.Y && m[down] != '#' && !slices.Contains(path, down) {
		options = append(options, down)
	}
	left := pos.Left()
	if left.X >= 0 && m[left] != '#' && !slices.Contains(path, left) {
		options = append(options, left)
	}
	right := pos.Right()
	if right.X < dim.X && m[right] != '#' && !slices.Contains(path, right) {
		options = append(options, right)
	}
	return options
}
