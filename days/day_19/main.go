package day_19

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"maps"
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
	workflows, _ := parseInput(input)
	sum := execCount("in", map[string][]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}, workflows)
	return strconv.Itoa(sum)
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
			if (matchesCondition[2] == "<" && ratings[matchesCondition[1]] < utils.Atoi(matchesCondition[3])) || (matchesCondition[2] == ">" && ratings[matchesCondition[1]] > utils.Atoi(matchesCondition[3])) {
				return exec(matchesCondition[4], ratings, workflows)
			}
		} else {
			return exec(step, ratings, workflows)
		}
	}
	return true
}

// execCount counts how many different rating combinations will yield A starting from the current workflow
func execCount(fn string, ratings map[string][]int, workflows map[string][]string) int {
	if fn == "A" {
		return multiplyAll(ratings)
	} else if fn == "R" {
		return 0
	}
	counts := 0
	workflow := workflows[fn]
	reCondition := regexp.MustCompile(`(?P<a>\w)(?P<b>[<>])(?P<c>\d+):(?P<d>\w+)`)
	for _, step := range workflow {
		matchesCondition := reCondition.FindStringSubmatch(step)
		if len(matchesCondition) > 0 {
			clone := maps.Clone(ratings)
			if matchesCondition[2] == "<" {
				clone[matchesCondition[1]] = []int{clone[matchesCondition[1]][0], min(clone[matchesCondition[1]][1], utils.Atoi(matchesCondition[3])-1)}
				counts += execCount(matchesCondition[4], clone, workflows)
				ratings[matchesCondition[1]] = []int{max(ratings[matchesCondition[1]][0], utils.Atoi(matchesCondition[3])), ratings[matchesCondition[1]][1]}
			} else {
				clone[matchesCondition[1]] = []int{max(clone[matchesCondition[1]][0], utils.Atoi(matchesCondition[3])+1), clone[matchesCondition[1]][1]}
				counts += execCount(matchesCondition[4], clone, workflows)
				ratings[matchesCondition[1]] = []int{ratings[matchesCondition[1]][0], min(ratings[matchesCondition[1]][1], utils.Atoi(matchesCondition[3]))}
			}
		} else {
			counts += execCount(step, ratings, workflows)
		}
	}
	return counts
}

// addAll adds all ratings of a given rating map
func addAll(rating map[string]int) int {
	sum := 0
	for _, i := range rating {
		sum += i
	}
	return sum
}

// multiplyAll multiplies all possible rating combinations
func multiplyAll(ratings map[string][]int) int {
	product := 1
	for _, ints := range ratings {
		product *= max(ints[1]-ints[0]+1, 0)
	}
	return product
}
