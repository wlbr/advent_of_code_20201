package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdata(input string) (fishes []*Fish) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for _, line := range lines {
		numstrings := strings.Split(line, ",")
		for _, num := range numstrings {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error converting '%s' to int:  %s", num, err)
			}
			fishes = append(fishes, NewFish(n))
		}
	}
	return fishes
}
