package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readdata(input string) (levels [][]*octopus) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var intline []*octopus
		for x, c := range line {
			if l, e := strconv.Atoi(string(c)); e != nil {
				log.Fatalf("Unknown char '%c' in line [%s]\n", c, line)
			} else {
				intline = append(intline, &octopus{level: l, x: x, y: y})
			}
		}
		y++
		levels = append(levels, intline)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return levels
}
