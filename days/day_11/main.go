package day_11

import (
	"fmt"
	"math"
	"strconv"
)

// coordinates define a pair of X Y values indicating the position on a 2D map
type coordinates struct {
	x int
	y int
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
	sum := 0
	g := findGalaxies(input, 2)
	for _, g1 := range g {
		for _, g2 := range g {
			dx := float64(g2.x - g1.x)
			dy := float64(g2.y - g1.y)
			sum += int(math.Abs(dx) + math.Abs(dy))
		}
	}
	return strconv.Itoa(sum / 2)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	g := findGalaxies(input, 1000000)
	for _, g1 := range g {
		for _, g2 := range g {
			dx := float64(g2.x - g1.x)
			dy := float64(g2.y - g1.y)
			sum += int(math.Abs(dx) + math.Abs(dy))
		}
	}
	return strconv.Itoa(sum / 2)
}

// findGalaxies iterates over the input and finds the corrected location of each galaxy
func findGalaxies(input []string, o int) []coordinates {
	var offset coordinates
	var g []coordinates
	for y, row := range input {
		if rowEmpty(input, y) {
			offset.y += o - 1
		}
		offset.x = 0
		for x, c := range row {
			if columnEmpty(input, x) {
				offset.x += o - 1
			}
			if c == '#' {
				g = append(g, coordinates{x + offset.x, y + offset.y})
			}
		}
	}
	return g
}

// columnEmpty checks whether the given column of the input contains no galaxies
func columnEmpty(input []string, x int) bool {
	for y := 0; y < len(input); y++ {
		if input[y][x] == '#' {
			return false
		}
	}
	return true
}

// rowEmpty checks whether the given row of the input contains no galaxies
func rowEmpty(input []string, y int) bool {
	for _, c := range input[y] {
		if c == '#' {
			return false
		}
	}
	return true
}
