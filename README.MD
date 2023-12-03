# Advent of Code 2023

A lightweight execution environment for the [advent of code](https://adventofcode.com/) challenges written in Go.

## Get started

In the following section, I'll show you how to set up and run the environment for the challenges.

### Prerequisites

Make sure you have Go >= 1.21.4 installed.

### Clone the repository

```sh
git clone https://github.com/wlchs/aoc23
```

### Test the solutions

Before submitting a solution, you can test the algorithms with the template input. Navigate to the day you want to check and run the
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