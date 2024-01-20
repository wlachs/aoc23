package day_24

import (
	"fmt"
	"regexp"
	"strconv"
)

// FloatVec3 is the float64 implementation of the Vec3 type
type FloatVec3 struct {
	X float64
	Y float64
	Z float64
}

// HailStone represents a pair of 3D vectors.
// The p0 vector holds the initial position of the hailstone, whereas the v vector holds its direction
type HailStone struct {
	p0 FloatVec3
	v  FloatVec3
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
	hailStones := readInput(input)
	rangeStart := float64(200000000000000)
	rangeEnd := float64(400000000000000)
	return strconv.Itoa(countIntersections(hailStones, rangeStart, rangeEnd))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// readInput reads the input and creates a slice of hailstones from it
func readInput(input []string) []HailStone {
	res := make([]HailStone, len(input))

	for i, line := range input {
		res[i] = readLine(line)
	}

	return res
}

// readLine parses an input line and creates a HailStone based on it
func readLine(line string) HailStone {
	re := regexp.MustCompile(`-*\d+`)
	match := re.FindAllString(line, -1)
	m := make([]float64, 6)
	for i, s := range match {
		f, _ := strconv.ParseFloat(s, 64)
		m[i] = f
	}

	return HailStone{
		p0: FloatVec3{
			X: m[0],
			Y: m[1],
			Z: m[2],
		},
		v: FloatVec3{
			X: m[3],
			Y: m[4],
			Z: m[5],
		},
	}
}

// countIntersections counts how many if the given hailstones intersect within the given coordinate range
func countIntersections(stones []HailStone, start float64, end float64) int {
	count := 0

	for i := range stones {
		for j := range stones[i+1:] {
			x, y, t1, t2, ok := doMatch2D(&stones[i], &stones[i+j+1])
			if ok && t1 >= 0 && t2 >= 0 && x >= start && y >= start && x <= end && y <= end {
				count++
			}
		}
	}

	return count
}

// doMatch2D checks at which 2D location do the two given hailstones intersect with each other
func doMatch2D(a *HailStone, b *HailStone) (float64, float64, float64, float64, bool) {
	m1 := a.v.Y / a.v.X
	m2 := b.v.Y / b.v.X
	x1 := a.p0.X
	y1 := a.p0.Y
	x2 := b.p0.X
	y2 := b.p0.Y

	if m1 == m2 {
		return 0, 0, 0, 0, false
	}

	x := ((m1*x1 - y1) - (m2*x2 - y2)) / (m1 - m2)
	y := m1*(x-x1) + y1
	t1 := (x - x1) / a.v.X
	t2 := (x - x2) / b.v.X

	return x, y, t1, t2, true
}
