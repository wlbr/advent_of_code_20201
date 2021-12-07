package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readdata(input string) (depths []int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if d, err := strconv.Atoi(line); err == nil {
			depths = append(depths, d)
		} else {
			log.Fatalf("Error parsing dataset '%s':  %s", line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return depths
}
