package day_04

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
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
	sum := 0
	for _, card := range input {
		sum += calculateCardValue(card)
	}
	return strconv.Itoa(sum)
}

// calculateCardValue calculates how much a card is worth based on the number of winning cards you have
func calculateCardValue(card string) int {
	w, a := parseCard(card)
	c := matchCount(w, a)
	return int(math.Pow(2, float64(c-1)))
}

// parseCard reads a card from the input and returns the list of winning numbers and the list of all numbers
func parseCard(card string) ([]int, []int) {
	re := regexp.MustCompile("^Card\\s+\\d+: (?P<w>.+) \\| (?P<c>.+)$")
	match := re.FindStringSubmatch(card)
	splitter := regexp.MustCompile("\\s+")
	winningNumbers := splitter.Split(strings.TrimSpace(match[1]), -1)
	allNumbers := splitter.Split(strings.TrimSpace(match[2]), -1)
	return toIntSlice(winningNumbers), toIntSlice(allNumbers)
}

// toIntSlice converts a string slice to an int slice
func toIntSlice(numbers []string) []int {
	var s []int
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		s = append(s, n)
	}
	return s
}

// matchCount counts the numbers which can be found in both input slices
func matchCount(a []int, b []int) int {
	c := 0
	for _, i := range a {
		if slices.Contains(b, i) {
			c++
		}
	}
	return c
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
