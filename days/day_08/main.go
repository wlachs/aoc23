package day_08

import (
	"fmt"
	"regexp"
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
	instructions, nodes := getNodes(input)
	count := countStepsFromAToZ("AAA", nodes, instructions)
	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	instructions, nodes := getNodes(input)
	count := countGhostSteps(nodes, instructions)
	return strconv.Itoa(count)
}

// getNodes retrieves the network from the input and returns a graph as a map.
// The key of the map is the node and the value contains the edges.
func getNodes(input []string) (string, map[string][]string) {
	re := regexp.MustCompile("\\w{3}")
	goal := input[0]
	m := map[string][]string{}
	for _, s := range input[2:] {
		match := re.FindAllString(s, -1)
		m[match[0]] = []string{match[1], match[2]}
	}
	return goal, m
}

// countStepsFromAToZ start iterating over the input instructions starting from a node ending with "A".
// The iteration stops when a "Z" node is reached
func countStepsFromAToZ(current string, nodes map[string][]string, instructions string) int {
	steps := 0
	for ; !strings.HasSuffix(current, "Z"); steps++ {
		instruction := instructions[steps%len(instructions)]
		if instruction == 'L' {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}
	}
	return steps
}

// countGhostSteps start iterating over the input instructions starting from every input ending with "A".
// The iteration stops when every parallel iteration stands on a node ending with "Z"
func countGhostSteps(nodes map[string][]string, instructions string) int {
	var current []string
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			current = append(current, node)
		}
	}
	steps := 1
	for _, c := range current {
		steps = lcm(steps, countStepsFromAToZ(c, nodes, instructions))
	}
	return steps
}

// gcd calculates the greatest common divisor of a and b.
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lcm calculates the least common multiple of a and b.
func lcm(a int, b int) int {
	return (a / gcd(a, b)) * b
}
