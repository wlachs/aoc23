package day_23

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"slices"
	"strconv"
)

// Node represents a node of the hiking graph
type Node struct {
	pos   types.Vec2
	edges []Edge
}

// Edge represents a weighted connection between two hiking nodes
type Edge struct {
	target   *Node
	distance int
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
	start, end := buildGraph(m, false)
	return strconv.Itoa(findLongestPath([]Edge{start.edges[0]}, end))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	start, end := buildGraph(m, true)
	return strconv.Itoa(findLongestPath([]Edge{start.edges[0]}, end))
}

// buildGraph iterates over the map and builds a graph where the edge weights correspond to the distances between neighbouring junctions
func buildGraph(m map[types.Vec2]int32, ignoreSlopes bool) (*Node, *Node) {
	root := &Node{pos: findStart(m)}
	nodeMap := map[types.Vec2]*Node{}
	nodeMap[root.pos] = root
	nodes := []*Node{root}
	dim := bottomRight(m)

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		edges := findEdges(m, node, nodeMap, ignoreSlopes, dim)

		for _, edge := range edges {
			containsEdge := slices.ContainsFunc(node.edges, func(e Edge) bool {
				return e.target == edge.target && e.distance == edge.distance
			})

			if !containsEdge {
				nodes = append(nodes, edge.target)
				node.edges = append(node.edges, edge)
			}
		}
	}

	return root, nodeMap[findEnd(m)]
}

// findStart finds the starting node of the graph
func findStart(m map[types.Vec2]int32) types.Vec2 {
	for x := 0; ; x++ {
		vec := types.Vec2{X: x}
		if m[vec] == '.' {
			return vec
		}
	}
}

// findEnd finds the final node of the graph
func findEnd(m map[types.Vec2]int32) types.Vec2 {
	dim := bottomRight(m)
	for x := 0; ; x++ {
		vec := types.Vec2{X: x, Y: dim.Y - 1}
		if m[vec] == '.' {
			return vec
		}
	}
}

// bottomRight finds the element at the bottom right position of the map
func bottomRight(m map[types.Vec2]int32) types.Vec2 {
	r := types.Vec2{X: 0, Y: 0}
	for vec := range m {
		r.X = max(r.X, vec.X+1)
		r.Y = max(r.Y, vec.Y+1)
	}
	return r
}

// findEdges finds the neighbouring nodes and their distances
func findEdges(m map[types.Vec2]int32, node *Node, nodeMap map[types.Vec2]*Node, ignoreSlopes bool, dim types.Vec2) []Edge {
	e := make([]Edge, 0, 4)
	initialOptions := findNextOptions(m, []types.Vec2{node.pos}, dim, ignoreSlopes)

	for _, option := range initialOptions {
		path := []types.Vec2{node.pos, option}
		options := []types.Vec2{option}

		for len(options) == 1 {
			path = append(path, options[0])
			options = findNextOptions(m, path, dim, ignoreSlopes)
		}

		junction := path[len(path)-1]
		n, ok := nodeMap[junction]
		if !ok {
			n = &Node{pos: junction}
			nodeMap[junction] = n
		}

		e = append(e, Edge{target: n, distance: len(path) - 2})
	}

	return e
}

// findNextOptions finds the next possible junction on the hiking path
func findNextOptions(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2, ignoreSlopes bool) []types.Vec2 {
	if ignoreSlopes {
		return findNonSlipperyNextOptions(m, path, dim)
	} else {
		return findSlipperyNextOptions(m, path, dim)
	}
}

// findSlipperyNextOptions checks the surrounding fields for the next step of the hike
func findSlipperyNextOptions(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2) []types.Vec2 {
	options := make([]types.Vec2, 0, 4)
	pos := path[len(path)-1]
	up := pos.Up()
	if up.Y >= 0 && m[up] != 'v' && m[up] != '#' && !slices.Contains(path, up) {
		options = append(options, up)
	}
	down := pos.Down()
	if down.Y < dim.Y && m[down] != '^' && m[down] != '#' && !slices.Contains(path, down) {
		options = append(options, down)
	}
	left := pos.Left()
	if left.X >= 0 && m[left] != '>' && m[left] != '#' && !slices.Contains(path, left) {
		options = append(options, left)
	}
	right := pos.Right()
	if right.X < dim.X && m[right] != '<' && m[right] != '#' && !slices.Contains(path, right) {
		options = append(options, right)
	}
	return options
}

// findNonSlipperyNextOptions checks the surrounding fields for the next step of the hike
// this variant of the method ignores steep slopes
func findNonSlipperyNextOptions(m map[types.Vec2]int32, path []types.Vec2, dim types.Vec2) []types.Vec2 {
	options := make([]types.Vec2, 0, 4)
	pos := path[len(path)-1]
	up := pos.Up()
	if up.Y >= 0 && m[up] != '#' && !slices.Contains(path, up) {
		options = append(options, up)
	}
	down := pos.Down()
	if down.Y < dim.Y && m[down] != '#' && !slices.Contains(path, down) {
		options = append(options, down)
	}
	left := pos.Left()
	if left.X >= 0 && m[left] != '#' && !slices.Contains(path, left) {
		options = append(options, left)
	}
	right := pos.Right()
	if right.X < dim.X && m[right] != '#' && !slices.Contains(path, right) {
		options = append(options, right)
	}
	return options
}

// findLongestPath iterates over the input map and finds the longest hiking path without loops
func findLongestPath(path []Edge, end *Node) int {
	pathLength := len(path)
	node := path[pathLength-1].target

	if node == end {
		sum := 0
		for _, edge := range path {
			sum += edge.distance
		}
		return sum
	}

	longest := 0
	for _, edge := range node.edges {
		containsNode := slices.ContainsFunc(path, func(e Edge) bool {
			return e.target == edge.target
		})

		if !containsNode {
			newPath := slices.Clone(path)
			newPath = append(newPath, edge)
			longest = max(longest, findLongestPath(newPath, end))
		}
	}
	return longest
}
