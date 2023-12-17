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
	path   []types.Vec2
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
	findShortestPath(m, mem, &br)
	return strconv.Itoa(mem[br])
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
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

// findShortestPath navigates through the map while trying to find the path from top left to bottom right with the minimal possible heat loss
func findShortestPath(m map[types.Vec2]int32, mem map[types.Vec2]int, br *types.Vec2) {
	nodes := map[PosDirRem]int{
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{X: 1}, rem: 3}}: mem[types.Vec2{}],
	}
	history := map[types.Vec2][]DirRem{}
	for len(nodes) > 0 {
		current := getNextNode(nodes)
		delete(nodes, current.PosDirRem)
		history[current.pos] = append(history[current.pos], current.DirRem)
		neighbours := getNeighbours(m, mem, current, br)
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

// getNeighbours gets the shortest path from the current node to its neighbours
func getNeighbours(m map[types.Vec2]int32, mem map[types.Vec2]int, node *Record, br *types.Vec2) []Record {
	var res []Record
	l := node.pos.Around()
	for _, v := range l {
		if node.pos.Subtract(&v) != node.dir {
			dir := v.Subtract(&node.pos)
			rem := 2
			if dir == node.dir {
				rem = node.rem - 1
			}
			if rem >= 0 && inBounds(&v, br) {
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
					path:   slices.Clone(node.path),
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

func printSteps(mem map[types.Vec2]int) {
	for y := 0; y <= 12; y++ {
		fmt.Print("| ")
		for x := 0; x <= 12; x++ {
			fmt.Printf(" %03d |", mem[types.Vec2{X: x, Y: y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

// inBounds makes sure that the current location is within the map coordinates
func inBounds(cur *types.Vec2, dimensions *types.Vec2) bool {
	return !(cur.X < 0 || cur.Y < 0 || cur.X > dimensions.X || cur.Y > dimensions.Y)
}
