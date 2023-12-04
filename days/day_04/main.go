package day_04

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// card struct representing a scratchcard
type card struct {
	index   int
	winners []int
	all     []int
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
	for _, c := range input {
		sum += calculateCardValue(c)
	}
	return strconv.Itoa(sum)
}

// calculateCardValue calculates how much a card is worth based on the number of winning cards you have
func calculateCardValue(card string) int {
	c := parseCard(card)
	return int(math.Pow(2, float64(c.matchCount()-1)))
}

// parseCard reads a card from the input and returns the list of winning numbers and the list of all numbers
func parseCard(c string) card {
	re := regexp.MustCompile("^Card\\s+(?P<i>\\d+): (?P<w>.+) \\| (?P<c>.+)$")
	match := re.FindStringSubmatch(c)
	splitter := regexp.MustCompile("\\s+")
	index := match[1]
	i, _ := strconv.Atoi(index)
	winningNumbers := splitter.Split(strings.TrimSpace(match[2]), -1)
	allNumbers := splitter.Split(strings.TrimSpace(match[3]), -1)
	return card{i, toIntSlice(winningNumbers), toIntSlice(allNumbers)}
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
func (c card) matchCount() int {
	count := 0
	for _, i := range c.winners {
		if slices.Contains(c.all, i) {
			count++
		}
	}
	return count
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return strconv.Itoa(calculateCardCount(input))
}

// calculateCardCount calculates how many cards there will be at the end
func calculateCardCount(input []string) int {
	var cards []card
	cardCount := map[int]int{}
	for _, c := range input {
		p := parseCard(c)
		cards = append(cards, p)
		cardCount[p.index] = 1
	}
	for _, c := range cards {
		toAdd := c.matchCount()
		for idToAdd := c.index + 1; idToAdd <= c.index+toAdd; idToAdd++ {
			_, canAdd := cardCount[idToAdd]
			_, hasCount := cardCount[c.index]
			if canAdd && hasCount {
				cardCount[idToAdd] += cardCount[c.index]
			}
		}
	}
	count := 0
	for _, i := range cardCount {
		count += i
	}
	return count
}
