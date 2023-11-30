package internal

import (
	"bufio"
	"fmt"
	"os"
)

// LoadInputLines loads a text file from the given path as a string slice.
func LoadInputLines(path string) []string {
	if path == "" {
		fmt.Println("missing input path")
		os.Exit(1)
	}

	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file \"%s\"\n", path)
		}
	}(file)

	if err != nil {
		fmt.Printf("failed to load input from \"%s\"\n", path)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to read file, %s\n", err)
	}

	return lines
}

// LoadFirstInputLine loads the first line of a text file from the given path as a string.
func LoadFirstInputLine(path string) string {
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file \"%s\"\n", path)
		}
	}(file)

	if err != nil {
		fmt.Printf("failed to load input from \"%s\"\n", path)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
