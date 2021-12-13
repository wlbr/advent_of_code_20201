package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdot(dotline string) (x, y int) {
	fmt.Sscanf(dotline, "%d,%d", &x, &y)
	return x, y
}

func readops(opslines []string) (ops []*instruction) {
	for _, line := range opslines {
		dir := ""
		parts := strings.Split(line, "=")
		if strings.HasPrefix(parts[0], "fold along ") {
			dir = string(parts[0][11])
		}
		op, _ := strconv.Atoi(parts[1])
		ops = append(ops, &instruction{pos: op, vertical: dir == "y"})
	}
	return ops
}

func getDims(dots []string) (dimx, dimy int) {
	dimx, dimy = readdot(dots[0])

	for _, dot := range dots[1:] {
		x, y := readdot(dot)

		if x > dimx {
			dimx = x
		}
		if y > dimy {
			dimy = y
		}
	}
	return dimx + 1, dimy + 1
}

func newDots(dotlines []string) *dots {
	dimx, dimy := getDims(dotlines)
	dots := &dots{startx: 0, starty: 0, endx: dimx, endy: dimy}

	for y := 0; y < dimy; y++ {
		dots.points = append(dots.points, make([]bool, dimx))
	}

	for _, line := range dotlines {
		x, y := readdot(line)
		dots.points[y][x] = true
	}
	return dots
}

func readdata(input string) (dots *dots, ops []*instruction) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dotlines, opslines []string
	readingops := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingops = true
		} else {
			if !readingops {
				dotlines = append(dotlines, line)
			} else {
				opslines = append(opslines, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	dots = newDots(dotlines)
	ops = readops(opslines)

	return dots, ops
}
