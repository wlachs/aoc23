package day_19

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"regexp"
	"strconv"
	"strings"
)

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
	workflows, ratings := parseInput(input)
	sum := 0
	for _, rating := range ratings {
		if exec("in", rating, workflows) {
			sum += addAll(rating)
		}
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// parseInput processes workflows and ratings of the input
func parseInput(input []string) (map[string][]string, []map[string]int) {
	reWorkflow := regexp.MustCompile(`(?P<a>\w+){(?P<b>.+)}`)
	reRating := regexp.MustCompile(`(?P<a>\w)=(?P<b>\d+)`)
	workflows := map[string][]string{}
	var ratings []map[string]int
	for _, row := range input {
		matchesWorkflow := reWorkflow.FindStringSubmatch(row)
		matchesRating := reRating.FindAllStringSubmatch(row, -1)
		if len(matchesWorkflow) > 0 {
			parseWorkflow(matchesWorkflow[1:], workflows)
		} else if len(matchesRating) > 0 {
			ratings = append(ratings, parseRating(matchesRating))
		}
	}
	return workflows, ratings
}

// parseWorkflow reads a single workflow and adds it to the workflow map passed as input
func parseWorkflow(s []string, workflows map[string][]string) {
	workflows[s[0]] = strings.Split(s[1], ",")
}

// parseRating reads a single line of ratings and creates a rating map from them
func parseRating(r [][]string) map[string]int {
	ratings := map[string]int{}
	for _, rating := range r {
		ratings[rating[1]] = utils.Atoi(rating[2])
	}
	return ratings
}

// exec executes a workflow with the given ratings
func exec(fn string, ratings map[string]int, workflows map[string][]string) bool {
	if fn == "A" {
		return true
	} else if fn == "R" {
		return false
	}
	workflow := workflows[fn]
	reCondition := regexp.MustCompile(`(?P<a>\w)(?P<b>[<>])(?P<c>\d+):(?P<d>\w+)`)
	for _, step := range workflow {
		matchesCondition := reCondition.FindStringSubmatch(step)
		if len(matchesCondition) > 0 {
			if matchesCondition[2] == "<" && ratings[matchesCondition[1]] < utils.Atoi(matchesCondition[3]) {
				return exec(matchesCondition[4], ratings, workflows)
			} else if matchesCondition[2] == ">" && ratings[matchesCondition[1]] > utils.Atoi(matchesCondition[3]) {
				return exec(matchesCondition[4], ratings, workflows)
			}
		} else {
			return exec(step, ratings, workflows)
		}
	}
	return true
}

// addAll adds all ratings of a given rating map
func addAll(rating map[string]int) int {
	sum := 0
	for _, i := range rating {
		sum += i
	}
	return sum
}
