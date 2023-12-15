package day_15

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"regexp"
	"strconv"
	"strings"
)

// lens consists of a string label and a focal length
type lens struct {
	label string
	focal int
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
	sum := 0
	for _, s := range strings.Split(input[0], ",") {
		sum += int(hash(s))
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	boxes := make([][]lens, 256)
	for _, s := range strings.Split(input[0], ",") {
		process(boxes, s)
	}
	sum := 0
	for boxID, box := range boxes {
		for lensID, l := range box {
			sum += (boxID + 1) * (lensID + 1) * l.focal
		}
	}
	return strconv.Itoa(sum)
}

// hash calculates the hash of the given string
// The hash is a numerical value from 0 to 255
func hash(s string) uint8 {
	h := uint8(0)
	for _, c := range s {
		h = 17 * (h + uint8(c))
	}
	return h
}

// process handles the given instruction and makes changes on the box slice based on it
func process(boxes [][]lens, i string) {
	re := regexp.MustCompile(`(?P<a>\w+)(?P<b>[-=])(?P<c>\d*)`)
	match := re.FindStringSubmatch(i)
	if match[2] == "=" {
		addLensToBox(&lens{label: match[1], focal: utils.Atoi(match[3])}, boxes)
	} else {
		removeLensFromBox(&lens{label: match[1]}, boxes)
	}
}

// addLensToBox adds (or replaces) the given lens to one of the boxes
func addLensToBox(l *lens, boxes [][]lens) {
	for i := range boxes[hash(l.label)] {
		if boxes[hash(l.label)][i].label == l.label {
			boxes[hash(l.label)][i].focal = l.focal
			return
		}
	}
	boxes[hash(l.label)] = append(boxes[hash(l.label)], *l)
}

// removeLensFromBox removes the lens with the given label from the boxes
func removeLensFromBox(l *lens, boxes [][]lens) {
	for i := range boxes[hash(l.label)] {
		if boxes[hash(l.label)][i].label == l.label {
			boxes[hash(l.label)] = append(boxes[hash(l.label)][:i], boxes[hash(l.label)][i+1:]...)
			return
		}
	}
}
