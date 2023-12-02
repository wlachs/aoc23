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
	red := 12
	green := 13
	blue := 14
	for rowId, game := range input {
		if isPossible(game, red, green, blue) {
			sum += rowId + 1
		}
	}
	return strconv.Itoa(sum)
}

// isPossible checks whether the game with the given input is possible with the number of different cubes in the bag
func isPossible(game string, red int, green int, blue int) bool {
	sets := getSets(game)
	for _, s := range sets {
		if s.red > red || s.green > green || s.blue > blue {
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
	for _, set := range setStrings {
		sets = append(sets, getSet(set))
	}
	return sets
}

// getSet receives a set as string and creates a set struct from it
func getSet(setString string) set {
	var set set
	cubes := strings.Split(setString, ", ")
	for _, cube := range cubes {
		count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		if strings.HasSuffix(cube, "red") {
			set.red = count
		} else if strings.HasSuffix(cube, "green") {
			set.green = count
		} else if strings.HasSuffix(cube, "blue") {
			set.blue = count
		}
	}
	return set
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
