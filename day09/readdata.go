package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdata(input string) (heights [][]int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		sheightsline := strings.Split(line, "")
		intheightline := make([]int, len(sheightsline))
		for i, v := range sheightsline {
			intheightline[i], _ = strconv.Atoi(v)
		}

		heights = append(heights, intheightline)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return heights
}
