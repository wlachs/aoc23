package day_05

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// interval representing a range of numbers
type interval struct {
	source int
	length int
}

// startsIn checks whether the current interval starts within the provided one
func (i interval) startsIn(it interval) bool {
	return i.source >= it.source && i.source < it.source+it.length
}

// endsIn checks whether the current interval ends within the provided one
func (i interval) endsIn(it interval) bool {
	return i.source+i.length-1 >= it.source && i.source+i.length-1 < it.source+it.length
}

// contains checks whether the current interval fully contains the provided one
func (i interval) contains(it interval) bool {
	return i.source <= it.source && i.source+i.length >= it.source+it.length
}

// isInside checks whether the current interval is fully contained by the provided one
func (i interval) isInside(it interval) bool {
	return i.startsIn(it) && i.endsIn(it)
}

// intervalWithOffset representing a range of input to be mapped to a range of output
type intervalWithOffset struct {
	interval
	offset int
}

// doMap tries to map an input if it is in range
func (i intervalWithOffset) doMap(input int) (int, bool) {
	if input >= i.source && input < i.source+i.length {
		return input + i.offset, true
	}
	return input, false
}

// mappingLayer is responsible for holding a slice of intervals
type mappingLayer struct {
	intervals []intervalWithOffset
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

// reduce splits the input interval to a set of mapped output intervals
func (m mappingLayer) reduce(i interval) []interval {
	if i.length <= 0 {
		return []interval{}
	}
	current := interval{
		source: m.doMap(i.source),
		length: i.length,
	}
	for _, it := range m.intervals {
		if i.contains(it.interval) {
			if i.source == it.source {
				current.length = min(current.length, it.length)
			} else {
				current.length = min(current.length, it.source-i.source)
			}
		} else if i.isInside(it.interval) {
			break
		} else if i.startsIn(it.interval) {
			current.length = it.length + it.source - i.source
		} else if i.endsIn(it.interval) {
			current.length = it.source - i.source
		}
	}
	rest := m.reduce(interval{
		source: i.source + current.length,
		length: i.length - current.length,
	})
	return append(rest, current)
}

// reduceAll iterates over all input intervals and reduces it to their mapped sub-intervals
func (m mappingLayer) reduceAll(i []interval) []interval {
	var intervals []interval
	for _, it := range i {
		intervals = append(intervals, m.reduce(it)...)
	}
	return intervals
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

// reduceAll finds the seed with the smallest possible location for the given input intervals
func (w layerWrapper) reduceAll(i []interval) int {
	l := math.MaxInt
	for _, layer := range w.layers {
		i = layer.reduceAll(i)
	}
	for _, it := range i {
		l = min(l, it.source)
	}
	return l
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
	return strconv.Itoa(intervalChain.reduceAll(seeds))
}

// getInitialSeeds finds the initial seeds from the input
func getInitialSeeds(input []string) []interval {
	splitter := regexp.MustCompile("\\s+")
	seeds := splitter.Split(strings.Split(input[0], ": ")[1], -1)
	var s []interval
	for _, seed := range seeds {
		i, _ := strconv.Atoi(seed)
		s = append(s, interval{
			source: i,
			length: 1,
		})
	}
	return s
}

// getInitialSeedIntervals finds the initial seed intervals from the input
func getInitialSeedIntervals(input []string) []interval {
	splitter := regexp.MustCompile("\\s+")
	seeds := splitter.Split(strings.Split(input[0], ": ")[1], -1)
	var it []interval
	for i := 0; i < len(seeds); i += 2 {
		source, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])
		it = append(it, interval{
			source: source,
			length: length,
		})
	}
	return it
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
			layer.intervals = append(layer.intervals, intervalWithOffset{
				interval: interval{
					source: source,
					length: length,
				},
				offset: target - source,
			})
		}
	}
	layers = append(layers, layer)
	return layerWrapper{layers}
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	intervalChain := getIntervalMapping(input)
	seeds := getInitialSeedIntervals(input)
	return strconv.Itoa(intervalChain.reduceAll(seeds))
}
