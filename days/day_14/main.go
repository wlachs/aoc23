package day_14

import (
	"fmt"
	"strconv"
)

// coordinates define a pair of X Y values indicating the position on a 2D map
type coordinates struct {
	x int
	y int
}

// above gives back the coordinates directly above the current one
func (c coordinates) above() coordinates {
	return coordinates{c.x, c.y - 1}
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
	m := parse(input)
	roll(m)
	return strconv.Itoa(load(m))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// parse reads the input rows and puts the values into a map of coordinates
func parse(input []string) map[coordinates]int32 {
	m := map[coordinates]int32{}
	for y, row := range input {
		for x, c := range row {
			m[coordinates{x, y}] = c
		}
	}
	return m
}

// roll tries to roll every stone in the input as far north as possible
func roll(m map[coordinates]int32) {
	printMap(m)
	oldLoad := load(m)
	for coords, field := range m {
		above := coords.above()
		if field == 'O' && m[above] == '.' {
			m[coords] = '.'
			m[above] = 'O'
		}
	}
	newLoad := load(m)
	if oldLoad != newLoad {
		roll(m)
	}
}

// load calculates the overall load on the north support beams
func load(m map[coordinates]int32) int {
	c := corner(m)
	sum := 0
	for coords, field := range m {
		if field == 'O' {
			sum += c.y - coords.y + 1
		}
	}
	return sum
}

// corner finds the bottom right corner of the input map
func corner(m map[coordinates]int32) coordinates {
	c := coordinates{0, 0}
	for coords := range m {
		c.x = max(c.x, coords.x)
		c.y = max(c.y, coords.y)
	}
	return c
}

func printMap(m map[coordinates]int32) {
	c := corner(m)
	for y := 0; y <= c.y; y++ {
		for x := 0; x <= c.x; x++ {
			fmt.Print(string(m[coordinates{x, y}]))
		}
		fmt.Println()
	}
	fmt.Println()
}
