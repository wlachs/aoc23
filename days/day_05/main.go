package day_05

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// interval representing a range of input to be mapped to a range of output
type interval struct {
	source int
	offset int
	length int
}

// doMap tries to map an input if it is in range
func (i interval) doMap(input int) (int, bool) {
	if input >= i.source && input < i.source+i.length {
		return input + i.offset, true
	}
	return input, false
}

// mappingLayer is responsible for holding a slice of intervals
type mappingLayer struct {
	intervals []interval
}

// doMap tries to push an input through the intervals. If any matches, the mapped value will be returned, otherwise the input remains unchanged.
func (m mappingLayer) doMap(input int) int {
	for _, i := range m.intervals {
		v, ok := i.doMap(input)
		if ok {
			return v
		}
	}
	return input
}

// layerWrapper wraps all mapping layers into a single object for easier handling
type layerWrapper struct {
	layers []mappingLayer
}

// doMap maps the given seed to a location
func (w layerWrapper) doMap(input int) int {
	for _, layer := range w.layers {
		input = layer.doMap(input)
	}
	return input
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
	seeds := getInitialSeeds(input)
	intervalChain := getIntervalMapping(input)
	return strconv.Itoa(findLowestLocation(seeds, intervalChain))
}

// getInitialSeeds find the initial seeds from the input
func getInitialSeeds(input []string) []int {
	splitter := regexp.MustCompile("\\s+")
	seeds := splitter.Split(strings.Split(input[0], ": ")[1], -1)
	var s []int
	for _, seed := range seeds {
		i, _ := strconv.Atoi(seed)
		s = append(s, i)
	}
	return s
}

// getIntervalMapping returns the mapping layers
func getIntervalMapping(input []string) layerWrapper {
	var layers []mappingLayer
	var layer mappingLayer
	for _, s := range input[2:] {
		if s == "" {
			layers = append(layers, layer)
			layer = mappingLayer{}
		} else if !strings.Contains(s, "map") {
			splitter := regexp.MustCompile("\\s+")
			row := splitter.Split(s, -1)
			target, _ := strconv.Atoi(row[0])
			source, _ := strconv.Atoi(row[1])
			length, _ := strconv.Atoi(row[2])
			layer.intervals = append(layer.intervals, interval{
				source: source,
				offset: target - source,
				length: length,
			})
		}
	}
	layers = append(layers, layer)
	return layerWrapper{layers}
}

// findLowestLocation finds the lowest location number corresponding to any of the initial seeds
func findLowestLocation(seeds []int, layers layerWrapper) int {
	m := math.MaxInt
	for _, seed := range seeds {
		loc := layers.doMap(seed)
		if loc < m {
			m = loc
		}
	}
	return m
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
