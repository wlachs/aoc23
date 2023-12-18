package day_18

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"regexp"
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
	return strconv.Itoa(dig(i))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	i := readHexInput(input)
	return strconv.Itoa(dig(i))
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

// readHexInput reads the individual dig instruction from the input
func readHexInput(input []string) []DigInstruction {
	res := make([]DigInstruction, 0, len(input))
	re := regexp.MustCompile(`\w{6}`)
	for _, s := range input {
		match := re.FindString(s)
		var di DigInstruction
		switch match[5] {
		case '0':
			di.vec = types.Vec2{X: 1}
		case '1':
			di.vec = types.Vec2{Y: 1}
		case '2':
			di.vec = types.Vec2{X: -1}
		case '3':
			di.vec = types.Vec2{Y: -1}
		}
		decimal, _ := strconv.ParseInt(match[:5], 16, 32)
		di.length = int(decimal)
		res = append(res, di)
	}
	return res
}

// dig executes the given instructions and calculates the dig area
func dig(instructions []DigInstruction) int {
	cur := types.Vec2{}
	vertices := make([]types.Vec2, 0, len(instructions))
	walls := 0
	for _, instruction := range instructions {
		vertices = append(vertices, cur)
		delta := instruction.vec.Multiply(instruction.length)
		walls += instruction.length
		cur = cur.Add(&delta)
	}
	return polygonArea(vertices) + walls/2 + 1
}

// polygonArea calculates the area of any polygon
// https://mathopenref.com/coordpolygonarea2.html
func polygonArea(v []types.Vec2) int {
	area := 0
	j := len(v) - 1
	for i := 0; i < len(v); i++ {
		area += (v[i].X + v[j].X) * (v[i].Y - v[j].Y)
		j = i
	}
	return area / 2
}
