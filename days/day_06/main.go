package day_06

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"regexp"
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
	times, distances := getRaces(input)
	product := 1
	for i := 0; i < len(times); i++ {
		product *= countWaysToWin(times[i], distances[i])
	}
	return strconv.Itoa(product)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	time, distance := getRace(input)
	return strconv.Itoa(countWaysToWin(time, distance))
}

// getRaces reads the input and returns two slices.
// The first slice contains the duration of a race and the corresponding value in
// the second slice is the record to beat
func getRaces(input []string) ([]int, []int) {
	re := regexp.MustCompile("\\d+")
	times := re.FindAllString(input[0], -1)
	distances := re.FindAllString(input[1], -1)
	return utils.ToIntSlice(times), utils.ToIntSlice(distances)
}

// getRace reads the input race.
// The first int contains the duration of a race and the second one is the record to beat
func getRace(input []string) (int, int) {
	re := regexp.MustCompile("\\d+")
	time := ""
	for _, s := range re.FindAllString(input[0], -1) {
		time += s
	}
	distance := ""
	for _, s := range re.FindAllString(input[1], -1) {
		distance += s
	}
	return utils.Atoi(time), utils.Atoi(distance)
}

// countWaysToWin does as expected and counts the different possible ways to win a race.
// The first param indicates the remaining time and the second one the distance to beat.
func countWaysToWin(t int, d int) int {
	winCount := 0
	for holdingTime := 0; holdingTime <= t; holdingTime++ {
		if (t-holdingTime)*holdingTime > d {
			winCount++
		}
	}
	return winCount
}
