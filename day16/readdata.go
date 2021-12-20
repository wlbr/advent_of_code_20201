package main

import (
	"bufio"
	"log"
	"os"
)

func readdata(input string) (code string) {
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
		log.Fatal("reading standard input:", err)
	}
	if len(lines) != 1 {
		log.Fatalf("Expected exactly one line in dataset '%s'", input)
	}
	return lines[0]
}
