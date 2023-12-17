package day_17

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"math"
	"slices"
	"strconv"
)

// DirRem holds a record of the direction and remaining straight distance.
type DirRem struct {
	dir types.Vec2
	rem int
}

// PosDirRem holds a record of the position, direction and remaining straight distance.
type PosDirRem struct {
	pos types.Vec2
	DirRem
}

// Record holds a record of the list containing the shortest path to a given field.
type Record struct {
	PosDirRem
	weight int
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
	mem := map[types.Vec2]int{}
	br := bottomRight(m)
	findShortestPathV1(m, mem, &br)
	return strconv.Itoa(mem[br])
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	mem := map[types.Vec2]int{}
	br := bottomRight(m)
	findShortestPathV2(m, mem, &br)
	return strconv.Itoa(mem[br])
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

// findShortestPathV1 navigates through the map while trying to find the path from top left to bottom right with the minimal possible heat loss
// part 1
func findShortestPathV1(m map[types.Vec2]int32, mem map[types.Vec2]int, br *types.Vec2) {
	nodes := map[PosDirRem]int{
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{X: 1}, rem: 3}}: mem[types.Vec2{}],
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{Y: 1}, rem: 3}}: mem[types.Vec2{}],
	}
	history := map[types.Vec2][]DirRem{}
	for len(nodes) > 0 {
		current := getNextNode(nodes)
		delete(nodes, current.PosDirRem)
		history[current.pos] = append(history[current.pos], current.DirRem)
		neighbours := getNeighboursV1(m, mem, current, br)
		for _, record := range neighbours {
			if !slices.Contains(history[record.pos], record.DirRem) {
				w, ok := nodes[record.PosDirRem]
				if ok {
					nodes[record.PosDirRem] = min(w, record.weight)
				} else {
					nodes[record.PosDirRem] = record.weight
				}
			}
		}
	}
}

// findShortestPathV2 navigates through the map while trying to find the path from top left to bottom right with the minimal possible heat loss
// part 2
func findShortestPathV2(m map[types.Vec2]int32, mem map[types.Vec2]int, br *types.Vec2) {
	nodes := map[PosDirRem]int{
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{X: 1}, rem: 10}}: mem[types.Vec2{}],
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{Y: 1}, rem: 10}}: mem[types.Vec2{}],
	}
	history := map[types.Vec2][]DirRem{}
	for len(nodes) > 0 {
		current := getNextNode(nodes)
		delete(nodes, current.PosDirRem)
		history[current.pos] = append(history[current.pos], current.DirRem)
		neighbours := getNeighboursV2(m, mem, current, br)
		for _, record := range neighbours {
			if !slices.Contains(history[record.pos], record.DirRem) {
				w, ok := nodes[record.PosDirRem]
				if ok {
					nodes[record.PosDirRem] = min(w, record.weight)
				} else {
					nodes[record.PosDirRem] = record.weight
				}
			}
		}
	}
}

// getNextNode finds the node in the unvisited locations with the lowest weight function
func getNextNode(nodes map[PosDirRem]int) *Record {
	var smallestPDR PosDirRem
	smallestW := math.MaxInt
	for pdr, w := range nodes {
		if w < smallestW {
			smallestPDR = pdr
			smallestW = w
		}
	}
	return &Record{PosDirRem: smallestPDR, weight: smallestW}
}

// weight gets the loss value at a given specific location of the map
func weight(m map[types.Vec2]int32, pos *types.Vec2) int {
	return int(m[*pos] - '0')
}

// getNeighboursV1 gets the shortest path from the current node to its neighbours
// can only move 3 block straight in a row
func getNeighboursV1(m map[types.Vec2]int32, mem map[types.Vec2]int, node *Record, br *types.Vec2) []Record {
	var res []Record
	l := node.pos.Around()
	for _, v := range l {
		if node.pos.Subtract(&v) != node.dir {
			dir := v.Subtract(&node.pos)
			rem := 3
			if dir == node.dir {
				rem = node.rem - 1
			}
			if rem > 0 && inBounds(&v, br) {
				w := node.weight + weight(m, &v)
				oldW, ok := mem[v]
				rec := Record{
					PosDirRem: PosDirRem{
						pos: v,
						DirRem: DirRem{
							dir: dir,
							rem: rem,
						},
					},
					weight: w,
				}
				if ok {
					mem[v] = min(oldW, w)
				} else {
					mem[v] = w
				}
				res = append(res, rec)
			}
		}
	}
	return res
}

// getNeighboursV2 gets the shortest path from the current node to its neighbours
// must move at least 4 blocks straight but not more than 10
func getNeighboursV2(m map[types.Vec2]int32, mem map[types.Vec2]int, node *Record, br *types.Vec2) []Record {
	var res []Record
	var vectors []types.Vec2
	if node.rem > 7 {
		vectors = append(vectors, node.pos.Add(&node.dir))
	} else {
		vectors = node.pos.Around()
	}
	for _, v := range vectors {
		if node.pos.Subtract(&v) != node.dir {
			dir := v.Subtract(&node.pos)
			rem := 10
			if dir == node.dir {
				rem = node.rem - 1
			}
			if rem > 0 && inBounds(&v, br) && (v != *br || rem <= 7) {
				w := node.weight + weight(m, &v)
				oldW, ok := mem[v]
				rec := Record{
					PosDirRem: PosDirRem{
						pos: v,
						DirRem: DirRem{
							dir: dir,
							rem: rem,
						},
					},
					weight: w,
				}
				if ok {
					mem[v] = min(oldW, w)
				} else {
					mem[v] = w
				}
				res = append(res, rec)
			}
		}
	}
	return res
}

// inBounds makes sure that the current location is within the map coordinates
func inBounds(cur *types.Vec2, dimensions *types.Vec2) bool {
	return !(cur.X < 0 || cur.Y < 0 || cur.X > dimensions.X || cur.Y > dimensions.Y)
}
