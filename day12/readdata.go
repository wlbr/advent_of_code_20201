package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readdata(input string) *plan {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	allnodes := newPlan(nil, nil)

	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, "-")
		for _, name := range names {
			node := allnodes.add(name)
			switch name {
			case "start":
				allnodes.start = node
			case "end":
				allnodes.end = node
			}
		}
		//there are always two nodes
		allnodes.get(names[0]).addConnection(allnodes.get(names[1]))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return allnodes
}
