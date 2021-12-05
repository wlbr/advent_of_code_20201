package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readfile(input string) (lines []string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func safeAtoi(nums []string) (ints []int) {

	for _, n := range nums {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("Error converting command '%v' coords to int: %s", nums, err)
		}
		ints = append(ints, i)
	}
	return ints
}

var rex *regexp.Regexp = regexp.MustCompile(`^\s*(\d+),\s*(\d+)\s*->\s*(\d+),\s*(\d+)\s*$`)

func readdata(filename string) (program []*command) {
	lines := readfile(filename)

	for _, line := range lines {
		matches := rex.FindStringSubmatch(line)
		if len(matches) != 5 {
			log.Printf("Warning: Command not in standard format: '%s'", line)
		} else {
			coords := safeAtoi(matches[1:])
			program = append(program, &command{x1: coords[0], y1: coords[1], x2: coords[2], y2: coords[3]})
		}
	}
	return program
}
