package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) (polymers *node, rules rules) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	spolymers := scanner.Text()

	var last *node
	for i := len(spolymers) - 1; i >= 0; i-- {
		last = &node{polymer: rune(spolymers[i]), next: last}
	}
	polymers = last

	scanner.Scan()
	rules = make(map[rune]map[rune]rune)
	for scanner.Scan() {
		ruleline := scanner.Text()
		var fromL, fromR rune
		var to rune
		fmt.Sscanf(ruleline, "%c%c -> %c", &fromL, &fromR, &to)
		if left, ok := rules[fromL]; !ok {
			rules[fromL] = make(map[rune]rune)
			rules[fromL][fromR] = to
		} else {
			left[fromR] = to
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return polymers, rules
}
