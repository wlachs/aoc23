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
	m, s := readMap(input)
	return strconv.Itoa(getEnclosedPoints(&m, &s))
}

// dim holds the size of the input map
var dim coordinates

// readMap reads the pipes of the input into a map
func readMap(input []string) (map[coordinates]int32, coordinates) {
	m := map[coordinates]int32{}
	var s coordinates
	for y, row := range input {
		for x, c := range row {
			coords := coordinates{x, y}
			dim.x = max(dim.x, x)
			dim.y = max(dim.y, y)
			m[coords] = c
			if c == 'S' {
				s = coords
			}
		}
	}
	dim.x++
	dim.y++
	return m, s
}

// getFurthestPoint follows the pipes starting from the given coordinates and calculates the number of steps in which the furthest possible
// point can be reached
func getFurthestPoint(m *map[coordinates]int32, s *coordinates) int {
	for _, neighbour := range s.getNeighbours() {
		l, ok := followPipe(m, s, &neighbour)
		if ok {
			return len(l) / 2
		}
	}
	panic("no loop found!")
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

// getEnclosedPoints counts the number of ground (.) tiles enclosed by the pipe loop
func getEnclosedPoints(m *map[coordinates]int32, s *coordinates) int {
	for _, neighbour := range s.getNeighbours() {
		l, ok := followPipe(m, s, &neighbour)
		if ok {
			return getEnclosedPointsByLoop(m, l)
		}
	}
	panic("no loop found!")
}

// getEnclosedPointsByLoop gets the number points in the map enclosed by the provided loop
func getEnclosedPointsByLoop(m *map[coordinates]int32, loop []coordinates) int {
	loopLength := len(loop)
	for i := 0; i < loopLength; i++ {
		current := loop[i]
		next := loop[(i+1)%loopLength]
		previous := loop[((i-1%loopLength)+loopLength)%loopLength]
		nextDir := coordinates{next.x - current.x, next.y - current.y}
		previousDir := coordinates{current.x - previous.x, current.y - previous.y}
		markSides(m, loop, &current, &nextDir)
		markSides(m, loop, &current, &previousDir)
	}
	outside := (*m)[coordinates{-1, -1}]
	var inside int32
	if outside == 'A' {
		inside = 'B'
	} else {
		inside = 'A'
	}
	countInside := 0
	for _, tile := range *m {
		if tile == inside {
			countInside++
		}
	}
	return countInside
}

// markSides marks every node on each side of the current tile based on the orientation.
// Nodes on the left are marked with "A", nodes on the right with "B"
func markSides(m *map[coordinates]int32, loop []coordinates, current *coordinates, direction *coordinates) {
	switch *direction {
	case coordinates{1, 0}:
		mark(m, loop, &coordinates{current.x, current.y - 1}, 'A')
		mark(m, loop, &coordinates{current.x, current.y + 1}, 'B')
	case coordinates{0, -1}:
		mark(m, loop, &coordinates{current.x - 1, current.y}, 'A')
		mark(m, loop, &coordinates{current.x + 1, current.y}, 'B')
	case coordinates{-1, 0}:
		mark(m, loop, &coordinates{current.x, current.y + 1}, 'A')
		mark(m, loop, &coordinates{current.x, current.y - 1}, 'B')
	case coordinates{0, 1}:
		mark(m, loop, &coordinates{current.x + 1, current.y}, 'A')
		mark(m, loop, &coordinates{current.x - 1, current.y}, 'B')
	}
}

// mark recursively marks the fields reachable from the current position without running out of bounds encountering a pipe of the loop
func mark(m *map[coordinates]int32, loop []coordinates, c *coordinates, side int32) {
	if slices.Contains(loop, *c) || c.x < -1 || c.y < -1 || c.x > dim.x || c.y > dim.y || (*m)[*c] == side {
		return
	}
	(*m)[*c] = side
	for _, neighbor := range c.getNeighbours() {
		mark(m, loop, &neighbor, side)
	}
}
