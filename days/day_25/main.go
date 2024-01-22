package day_25

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

// Node represents a node of a graph
type Node struct {
	label          string
	connectedNodes []*Node
}

// Edge connecting two nodes
type Edge struct {
	u *Node
	v *Node
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	nodeMap := readGraph(input)

	for {
		cut, res := minCut(nodeMap)
		if cut == 6 {
			return strconv.Itoa(res)
		}
	}
}

// readGraph reads the input rows and constructs a graph from them
func readGraph(input []string) map[string]*Node {
	res := map[string]*Node{}
	for _, s := range input {
		firstSplit := strings.Split(s, ": ")
		label := firstSplit[0]
		secondSplit := strings.Split(firstSplit[1], " ")
		source, ok1 := res[label]
		if !ok1 {
			source = &Node{label: label}
			res[label] = source
		}
		for _, otherLabel := range secondSplit {
			target, ok2 := res[otherLabel]
			if !ok2 {
				target = &Node{label: otherLabel}
				res[otherLabel] = target
			}
			source.connectedNodes = append(source.connectedNodes, target)
			target.connectedNodes = append(target.connectedNodes, source)
		}
	}
	return res
}

// minCut calculates a possible min cut of the graph such that it is divided into two components
func minCut(nodeMap map[string]*Node) (int, int) {
	vertices := len(nodeMap)
	sets := make([][]*Node, 0, len(nodeMap))
	var edges []Edge
	for _, u := range nodeMap {
		sets = append(sets, []*Node{u})
		for _, v := range u.connectedNodes {
			edges = append(edges, Edge{u: u, v: v})
		}
	}

	for vertices > 2 {
		i := rand.Intn(len(edges))
		set1 := find(sets, edges[i].u)
		set2 := find(sets, edges[i].v)

		if set1 != set2 {
			vertices--
			sets = union(sets, set1, set2)
		}
	}

	cut := 0
	for _, edge := range edges {
		set1 := find(sets, edge.u)
		set2 := find(sets, edge.v)
		if set1 != set2 {
			cut++
		}
	}

	return cut, len(sets[0]) * len(sets[1])
}

// find finds the index of the set to which the u node belongs
func find(sets [][]*Node, u *Node) int {
	for i, set := range sets {
		if slices.Contains(set, u) {
			return i
		}
	}
	panic("node not found")
}

// union merges the sets with the given IDs and returns the new slice of sets
func union(sets [][]*Node, set1 int, set2 int) [][]*Node {
	sets[set1] = append(sets[set1], sets[set2]...)
	return append(sets[:set2], sets[set2+1:]...)
}
