package day_16

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"slices"
	"strconv"
)

// Beam represents a location and a direction vector of a beam.
type Beam struct {
	types.Vec2
	dir types.Vec2
}

// next calculates the next position(s) of the beam based on splitters and mirrors of the input
func (b Beam) next(m map[types.Vec2]int32) []Beam {
	nextLocation := b.Vec2.Add(&b.dir)
	var beams []Beam
	switch m[nextLocation] {
	case '.':
		beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir})
	case '/':
		if b.dir.X != 0 {
			b.dir = b.dir.RotateLeft()
		} else {
			b.dir = b.dir.RotateRight()
		}
		beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir})
	case '\\':
		if b.dir.Y == 0 {
			b.dir = b.dir.RotateRight()
		} else {
			b.dir = b.dir.RotateLeft()
		}
		beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir})
	case '-':
		if b.dir.X == 0 {
			beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir.RotateLeft()}, Beam{Vec2: nextLocation, dir: b.dir.RotateRight()})
		} else {
			beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir})
		}
	case '|':
		if b.dir.Y == 0 {
			beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir.RotateLeft()}, Beam{Vec2: nextLocation, dir: b.dir.RotateRight()})
		} else {
			beams = append(beams, Beam{Vec2: nextLocation, dir: b.dir})
		}
	}
	return beams
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
	m := utils.ParseInputToMap(input)
	v := map[types.Vec2]bool{}
	moveBeams(m, []Beam{
		{
			Vec2: types.Vec2{X: -1, Y: 0},
			dir:  types.Vec2{X: 1, Y: 0},
		},
	}, v)
	return strconv.Itoa(energy(v))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// moveBeams iterates over the map and marks visited fields by beams.
func moveBeams(m map[types.Vec2]int32, beams []Beam, v map[types.Vec2]bool) {
	var beamMemory []Beam
	for len(beams) > 0 {
		beamMemory = append(beamMemory, beams...)
		var newBeams []Beam
		for _, beam := range beams {
			for _, b := range beam.next(m) {
				if !slices.Contains(beamMemory, b) {
					newBeams = append(newBeams, b)
				}
			}
		}
		beams = newBeams
		for _, beam := range beams {
			v[beam.Vec2] = true
		}
	}
}

func energy(v map[types.Vec2]bool) int {
	sum := 0
	for _, val := range v {
		if val {
			sum++
		}
	}
	return sum
}

func printMap(m map[types.Vec2]bool) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			_, ok := m[types.Vec2{X: x, Y: y}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
