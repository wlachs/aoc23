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
	m := utils.ParseInputToMap(input)
	br := bottomRight(m)
	maxEnergy := 0
	for vec := range m {
		if vec.X > 0 && vec.X < br.X && vec.Y > 0 && vec.Y < br.Y {
			continue
		}
		v := map[types.Vec2]bool{}
		start := types.Vec2{X: 0, Y: 0}
		dir := types.Vec2{X: 0, Y: 0}
		if vec.X == 0 {
			start = types.Vec2{X: -1, Y: vec.Y}
			dir = types.Vec2{X: 1, Y: 0}
		} else if vec.X == br.X {
			start = types.Vec2{X: br.X + 1, Y: vec.Y}
			dir = types.Vec2{X: -1, Y: 0}
		} else if vec.Y == 0 {
			start = types.Vec2{X: vec.X, Y: -1}
			dir = types.Vec2{X: 0, Y: 1}
		} else if vec.Y == br.Y {
			start = types.Vec2{X: vec.X, Y: br.Y + 1}
			dir = types.Vec2{X: 0, Y: -1}
		}
		moveBeams(m, []Beam{
			{
				Vec2: start,
				dir:  dir,
			},
		}, v)
		maxEnergy = max(maxEnergy, energy(v))
	}
	return strconv.Itoa(maxEnergy)
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

// energy calculates the overall number of visited fields of the map
func energy(v map[types.Vec2]bool) int {
	sum := 0
	for _, val := range v {
		if val {
			sum++
		}
	}
	return sum
}

// bottomRight finds the element at the bottom right position of the map
func bottomRight(m map[types.Vec2]int32) types.Vec2 {
	r := types.Vec2{X: 0, Y: 0}
	for vec := range m {
		r.X = max(r.X, vec.X)
		r.Y = max(r.Y, vec.Y)
	}
	return r
}
