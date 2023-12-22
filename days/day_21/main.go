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
	m := utils.ParseInputToMap(input)
	s := findStart(m)
	return strconv.Itoa(countInfiniteFields(m, s))
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

// countInfiniteFields counts the number of reachable fields in 26501365 steps starting from the given coordinates.
// The map wraps around infinitely
func countInfiniteFields(m map[types.Vec2]int32, s types.Vec2) int {
	d := bottomRight(m)
	init := make([]int, d.X)
	delta := make([]int, d.X)
	prevs := make([]int, d.X)
	var accPrev []types.Vec2
	accNext := []types.Vec2{s}
	for n := 1; n < 3*d.X; n++ {
		var accDelta []types.Vec2
		for _, vec2 := range accNext {
			ns := infiniteNeighbours(m, vec2, d)
			for _, t := range ns {
				if !slices.Contains(accDelta, t) && !slices.Contains(accPrev, t) {
					accDelta = append(accDelta, t)
				}
			}
		}
		accPrev = accNext
		accNext = accDelta
		if n >= d.X && n < 2*d.X {
			prevs[n%d.X] = len(accNext)
		} else if n >= 2*d.X {
			delta[n%d.X] = len(accNext) - prevs[n%d.X]
		} else {
			init[n] = len(accNext)
		}
	}
	c := 0
	mx := 26501365
	for i := mx % 2; i <= mx; i += 2 {
		c += step(init, prevs, delta, i, d.X)
	}
	return c
}

// step counts the number of newly visited fields as the given iteration
func step(init []int, prevs []int, delta []int, x int, d int) int {
	if x < d {
		return init[x]
	}
	return prevs[x%d] + (x/d-1)*delta[x%d]
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

// infiniteNeighbours calculates the neighbouring vectors of the given one and checks whether the input map has a navigable field on them.
// The input map wraps around infinitely.
func infiniteNeighbours(m map[types.Vec2]int32, vec2 types.Vec2, d types.Vec2) []types.Vec2 {
	res := make([]types.Vec2, 0, 4)
	n := vec2.Around()
	for _, v := range n {
		vProj := types.Vec2{X: mod(v.X, d.X), Y: mod(v.Y, d.Y)}
		if m[vProj] == '.' || m[vProj] == 'S' {
			res = append(res, v)
		}
	}
	return res
}

// addIfNotAlready adds neighbouring elements to the slice if they are not already contained
func addIfNotAlready(acc []types.Vec2, neighbours []types.Vec2) []types.Vec2 {
	for _, neighbour := range neighbours {
		if !slices.Contains(acc, neighbour) {
			acc = append(acc, neighbour)
		}
	}
	return acc
}

// bottomRight finds the element at the bottom right position of the map
func bottomRight(m map[types.Vec2]int32) types.Vec2 {
	r := types.Vec2{X: 0, Y: 0}
	for vec := range m {
		r.X = max(r.X, vec.X+1)
		r.Y = max(r.Y, vec.Y+1)
	}
	return r
}

// mod implements a modulo function that returns an always positive remainder
func mod(i int, m int) int {
	return ((i % m) + m) % m
}
