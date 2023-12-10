package day_10

import (
	"fmt"
	"slices"
	"strconv"
)

var east = []int32{'-', 'L', 'F', 'S'}
var north = []int32{'|', 'L', 'J', 'S'}
var west = []int32{'-', 'J', '7', 'S'}
var south = []int32{'|', '7', 'F', 'S'}

// coordinates define a pair of X Y values indicating the position on a 2D map
type coordinates struct {
	x int
	y int
}

// getNeighbours gets the coordinates directly above, below, left and right to the current one
func (c coordinates) getNeighbours() []coordinates {
	return []coordinates{
		{c.x + 1, c.y},
		{c.x, c.y - 1},
		{c.x - 1, c.y},
		{c.x, c.y + 1},
	}
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
	m, s := readMap(input)
	return strconv.Itoa(getFurthestPoint(&m, &s))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// readMap reads the pipes of the input into a map
func readMap(input []string) (map[coordinates]int32, coordinates) {
	m := map[coordinates]int32{}
	var s coordinates
	for y, row := range input {
		for x, c := range row {
			coords := coordinates{x, y}
			m[coords] = c
			if c == 'S' {
				s = coords
			}
		}
	}
	return m, s
}

// getFurthestPoint follows the pipes starting from the given coordinates and calculates the number of steps in which the furthest possible
// point can be reached
func getFurthestPoint(m *map[coordinates]int32, s *coordinates) int {
	maxDistance := 0
	for _, neighbour := range s.getNeighbours() {
		l, ok := followPipe(m, s, &neighbour)
		if ok {
			maxDistance = max(maxDistance, len(l)/2)
		}
	}
	return maxDistance
}

// followPipe tries to follow the pipes of the input.
// The search terminates if:
// - the pipes no longer continue
// - a loop is encountered
func followPipe(m *map[coordinates]int32, previous *coordinates, current *coordinates) ([]coordinates, bool) {
	if !transitionAllowed(m, previous, current) {
		return nil, false
	}
	if (*m)[*current] == 'S' {
		return []coordinates{*current}, true
	}
	for _, neighbor := range current.getNeighbours() {
		if neighbor == *previous {
			continue
		}
		next, ok := followPipe(m, current, &neighbor)
		if ok {
			return append(next, *current), true
		}
	}
	return nil, false
}

// transitionAllowed checks whether a transition between the given two coordinates of the map is allowed or not
func transitionAllowed(m *map[coordinates]int32, s *coordinates, t *coordinates) bool {
	source := (*m)[*s]
	target := (*m)[*t]
	switch t.x - s.x {
	case 1:
		return slices.Contains(east, source) && slices.Contains(west, target)
	case -1:
		return slices.Contains(west, source) && slices.Contains(east, target)
	}
	switch t.y - s.y {
	case 1:
		return slices.Contains(south, source) && slices.Contains(north, target)
	case -1:
		return slices.Contains(north, source) && slices.Contains(south, target)
	}
	return false
}
