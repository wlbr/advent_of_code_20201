package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) (algo ImageEnhancerAlgorithm, iarr *InfiniteArray) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	for _, v := range line {
		algo = append(algo, v)
	}
	scanner.Scan()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	iarr = NewInfiniteArrayFromString(lines, '.')

	return algo, iarr
}
