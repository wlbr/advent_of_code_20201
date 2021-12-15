package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func safeAtoi(nums string, cave *Cave, y int) (levels []*Risklevel) {
	for x, n := range nums {
		i, err := strconv.Atoi(string(n))
		if err != nil {
			log.Fatalf("Error converting command '%v' coords to int: %s", nums, err)
		}
		levels = append(levels, &Risklevel{value: i, cave: cave, x: x, y: y})
	}
	return levels
}

func readdata(input string) *Cave {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cave := &Cave{}
	for scanner.Scan() {
		line := scanner.Text()
		cave.levels = append(cave.levels, safeAtoi(line, cave, cave.dimy))
		cave.dimy++
	}
	cave.dimx = len(cave.levels[0])

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return cave
}
