package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) (binaries []string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		binaries = append(binaries, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return binaries
}
