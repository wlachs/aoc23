# Advent of Code Project Template

A lightweight execution environment for the [advent of code](https://adventofcode.com/) challenges written in Go.

## Get started

In the following section, I'll show you how to set up and run the environment for the challenges.

### Prerequisites

Make sure you have Go >= 1.21.4 installed.

### Clone the repository

```sh
git clone https://github.com/wlchs/advent_of_code_go_template
```

### Solve the daily challenges.

I've generated a separate directory in the repository for each advent calendar day: from the 1st until the 24th. Additionally, the `day_0`
directory contains a solution for the first challenge from last year's competition.

To solve a daily challenge, navigate to the corresponding day and implement the solution. I've prepared two methods for each day, `Part1`
and`Part2`, indicating that every challenge has two parts. The return value of the methods are always strings. The methods receive the input
as a string slice, so you don't have to worry about reading anything from a file.

The solution for the challenge of December 1, 2022, looks like this:

```go

// getElfCalories reads the input lines and returns a slice of calories for each elf
func getElfCalories(input []string) []int {
    var calories []int
    currentElf := 0
    for _, i := range input {
        if i == "" {
            calories = append(calories, currentElf)
            currentElf = 0
        } else {
            cal, _ := strconv.Atoi(i)
            currentElf += cal
        }
    }
    calories = append(calories, currentElf)
    sort.Slice(calories, func (i, j int) bool {
        return calories[i] > calories[j]
    })
    return calories
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
    calories := getElfCalories(input)
    maxCalories := calories[0]
    return strconv.Itoa(maxCalories)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
    calories := getElfCalories(input)
    topCalories := 0
    for i := 0; i < 3; i++ {
        topCalories += calories[i]
    }
    return strconv.Itoa(topCalories)
}

```

Additionally, I've prepared unit tests for each day to try the solutions with a template input. The text files in the directories contain
the template input and output for both parts of the exercise. Paste the input to `input_test_x.txt` file and the expected solution
to`solution_x.txt`.

### Test your solution

Before submitting a solution, you can test your algorithm with the template input. Navigate to the day you want to check and run the
following command:

```sh
go test
```

Alternatively, you can also run the command from the project root.

```sh
go test ./days/day_xx/
```

### Compile and run

You can compile code and run with the actual input if all the tests pass. For this, first, run this command:

```sh
go build .
```

To run the solution, you need to provide a few extra arguments.
* the `--day` flag must be set to specify which day's solution should run
* the `--input` flag specifies the path of the file containing the actual input
* optionally, you can add the `--mode` flag to run one part of the daily challenge, accepted values are 1 and 2

And now the complete command:

```sh
./advent_of_code_go_template --day x --input path_to_input --mode 1
# or
go run . --day x --input path_to_input --mode 1
```

## Contribution

If you'd like to contribute to the project, open an issue or a pull request!