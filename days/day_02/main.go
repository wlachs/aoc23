package day_02

import (
	"fmt"
	"strconv"
	"strings"
)

// set represents a set of the game
type set struct {
	red   int
	green int
	blue  int
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
	s := set{
		red:   12,
		green: 13,
		blue:  14,
	}
	for rowId, game := range input {
		if isPossible(game, s) {
			sum += rowId + 1
		}
	}
	return strconv.Itoa(sum)
}

// isPossible checks whether the game with the given input is possible with the number of different cubes in the bag
func isPossible(game string, cubes set) bool {
	sets := getSets(game)
	for _, s := range sets {
		if s.red > cubes.red || s.green > cubes.green || s.blue > cubes.blue {
			return false
		}
	}
	return true
}

// getSets parses a single input game and creates a slice of sets out of it
func getSets(game string) []set {
	gameWithoutId := strings.Split(game, ": ")[1]
	setStrings := strings.Split(gameWithoutId, "; ")
	var sets []set
	for _, s := range setStrings {
		sets = append(sets, getSet(s))
	}
	return sets
}

// getSet receives a set as string and creates a set struct from it
func getSet(setString string) set {
	var s set
	cubes := strings.Split(setString, ", ")
	for _, cube := range cubes {
		count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		if strings.HasSuffix(cube, "red") {
			s.red = count
		} else if strings.HasSuffix(cube, "green") {
			s.green = count
		} else if strings.HasSuffix(cube, "blue") {
			s.blue = count
		}
	}
	return s
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, game := range input {
		sum += calculateMinPower(game)
	}
	return strconv.Itoa(sum)
}

// calculateMinPower calculates the power of the smallest set of cubes which still make the game possible
func calculateMinPower(game string) int {
	sets := getSets(game)
	cubes := calculateMinCubes(sets)
	return cubes.red * cubes.green * cubes.blue
}

// calculateMinCubes finds the smallest required value of each color to make the game possible.
// The smallest required value is always the largest number in the input for the respective color.
func calculateMinCubes(sets []set) set {
	var m set
	for _, s := range sets {
		if s.red > m.red {
			m.red = s.red
		}
		if s.green > m.green {
			m.green = s.green
		}
		if s.blue > m.blue {
			m.blue = s.blue
		}
	}
	return m
}
