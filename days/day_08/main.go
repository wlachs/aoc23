package day_08

import (
	"fmt"
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
	instructions, nodes := getNodes(input)
	count := countSteps(nodes, instructions)
	// findShortestRoutes("AAA", &nodes, &routes, 0)
	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// getNodes retrieves the network from the input and returns a graph as a map.
// The key of the map is the node and the value contains the edges.
func getNodes(input []string) (string, map[string][]string) {
	re := regexp.MustCompile("[A-Z]{3}")
	goal := input[0]
	m := map[string][]string{}
	for _, s := range input[2:] {
		match := re.FindAllString(s, -1)
		m[match[0]] = []string{match[1], match[2]}
	}
	return goal, m
}

// findShortestRoutes finds the shortest path to any given node starting from the given one.
func findShortestRoutes(root string, nodes *map[string][]string, routes *map[string]int, depth int) {
	best, ok := (*routes)[root]
	if ok && best < depth {
		return
	}
	(*routes)[root] = depth
	for _, n := range (*nodes)[root] {
		findShortestRoutes(n, nodes, routes, depth+1)
	}
}

// countSteps start iterating over the input instructions starting from "AAA".
// The iteration stops when the "ZZZ" node is reached
func countSteps(nodes map[string][]string, instructions string) int {
	current := "AAA"
	steps := 0
	for ; current != "ZZZ"; steps++ {
		instruction := instructions[steps%len(instructions)]
		if instruction == 'L' {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}
	}
	return steps
}
