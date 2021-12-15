package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) (polymers string, rules rules) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	polymers = scanner.Text()

	scanner.Scan()
	rules = make(map[string]string)
	for scanner.Scan() {
		ruleline := scanner.Text()

		var from, to string
		fmt.Sscanf(ruleline, "%s -> %s", &from, &to)
		if left, ok := rules[from]; !ok {
			rules[from] = to
		} else {
			log.Fatalf("Rule already exists for '%s->%s' : was '%s'", from, to, left)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return polymers, rules
}
