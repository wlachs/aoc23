package day_0

import (
	"fmt"
	"sort"
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

// getElfCalories reads the input lines and returns a slice of calories for each elf
func getElfCalories(input []string) []int {
	var calories []int
	currentElf := 0

	for _, i := range input {
		if i == "" {
			calories = append(calories, currentElf)
			currentElf = 0
		} else {
			cal, _ := strconv.Atoi(i)
			currentElf += cal
		}
	}
	calories = append(calories, currentElf)

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	calories := getElfCalories(input)
	maxCalories := calories[0]
	return strconv.Itoa(maxCalories)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	calories := getElfCalories(input)

	topCalories := 0
	for i := 0; i < 3; i++ {
		topCalories += calories[i]
	}

	return strconv.Itoa(topCalories)
}
