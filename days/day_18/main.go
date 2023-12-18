package day_18

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"strconv"
	"strings"
)

// DigInstruction contains a single row of input representing a digging vector
type DigInstruction struct {
	vec    types.Vec2
	length int
}

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
	i := readInput(input)
	m := dig(i)
	return strconv.Itoa(len(m))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// readInput reads the individual dig instruction from the input
func readInput(input []string) []DigInstruction {
	res := make([]DigInstruction, 0, len(input))
	for _, s := range input {
		i := strings.Split(s, " ")
		var di DigInstruction
		switch i[0] {
		case "U":
			di.vec = types.Vec2{Y: -1}
		case "D":
			di.vec = types.Vec2{Y: 1}
		case "L":
			di.vec = types.Vec2{X: -1}
		case "R":
			di.vec = types.Vec2{X: 1}
		}
		di.length = utils.Atoi(i[1])
		res = append(res, di)
	}
	return res
}

var mins = types.Vec2{}
var maxes = types.Vec2{}

// dig executes the given instructions and builds a dig map
func dig(instructions []DigInstruction) map[types.Vec2]string {
	cur := types.Vec2{}
	m := map[types.Vec2]string{}
	for _, instruction := range instructions {
		for i := 0; i < instruction.length; i++ {
			m[cur] = "#"
			cur = cur.Add(&instruction.vec)
		}
	}
	for vec := range m {
		mins.X = min(mins.X, vec.X)
		mins.Y = min(mins.Y, vec.Y)
		maxes.X = max(maxes.X, vec.X)
		maxes.Y = max(maxes.Y, vec.Y)
	}
	for _, instruction := range instructions {
		for i := 0; i < instruction.length; i++ {
			// printMap(m)
			r := instruction.vec.RotateRight()
			start := cur.Add(&r)
			floodFill(start, "#", m)
			cur = cur.Add(&instruction.vec)
		}
	}
	return m
}

func printMap(m map[types.Vec2]string) {
	for y := mins.Y; y <= maxes.Y; y++ {
		for x := mins.X; x <= maxes.X; x++ {
			f, ok := m[types.Vec2{X: x, Y: y}]
			if ok {
				fmt.Print(f)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// floodFill recursively fills the map with the given string
func floodFill(start types.Vec2, s string, m map[types.Vec2]string) {
	if canFill(start, m) {
		m[start] = s
		for _, newStart := range start.Around() {
			floodFill(newStart, s, m)
		}
	}
}

// canFill checks whether the given location vector is in bounds and can be filled by the flood fill algorithm
func canFill(pos types.Vec2, m map[types.Vec2]string) bool {
	_, ok := m[pos]
	return pos.X >= mins.X && pos.X <= maxes.X && pos.Y >= mins.Y && pos.Y <= maxes.Y && !ok
}
