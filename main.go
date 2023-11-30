package main

import (
	"flag"
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/internal"
	"os"
	"strconv"
)

// main entry point
// The --day parameter is required to choose which daily challenge should be executed.
// The --input parameter is also required, and it points to the input file that should be used for the challenge.
// The --mode parameter specifies which part of the challenge should be executed:
// - 1: only the first part
// - 2: only the second part
// - 3 or empty: both parts
func main() {
	d := flag.String("day", "", "day ID to execute")
	i := flag.String("input", "", "input file path")
	m := flag.String("mode", "3", "running mode")
	flag.Parse()

	if d == nil || i == nil {
		fmt.Println("missing required input params")
		os.Exit(1)
	}

	day, err := strconv.Atoi(*d)
	if err != nil {
		fmt.Println("couldn't parse day")
		os.Exit(1)
	}

	mode, err := strconv.Atoi(*m)
	if err != nil {
		fmt.Println("incorrect mode")
		os.Exit(1)
	}

	inputPath := *i
	internal.RunChallenge(day, inputPath, mode)
}
