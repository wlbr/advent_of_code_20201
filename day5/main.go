package main

import (
	"fmt"
)

type command struct {
	x1, y1, x2, y2 int
}

func (c *command) String() string {
	return fmt.Sprintf("%d,%x - %d,%d", c.x1, c.y1, c.x2, c.y2)
}

type diagram struct {
	lines [][]int
}

func maxPos(cmds []*command) (int, int) {
	var maxx, maxy int
	for _, cmd := range cmds {
		if cmd.x1 > maxx {
			maxx = cmd.x1
		}
		if cmd.x2 > maxx {
			maxx = cmd.x2
		}
		if cmd.y1 > maxy {
			maxy = cmd.y1
		}
		if cmd.y2 > maxy {
			maxy = cmd.y2
		}
	}
	return maxx, maxy
}

func NewDiagram(cmds []*command) *diagram {
	d := &diagram{}
	x, y := maxPos(cmds)
	d.lines = make([][]int, y+1)
	for i := range d.lines {
		d.lines[i] = make([]int, x+1)
	}
	return d
}

func (d *diagram) countOverlaps() int {
	var count int
	for _, line := range d.lines {
		for _, pos := range line {
			if pos > 1 {
				count++
			}
		}
	}
	return count
}

func step(x, y int, cmd *command) (int, int, bool) {
	if x == cmd.x2 && y == cmd.y2 {
		return -1, -1, true
	}

	if x > cmd.x2 {
		x -= 1
	} else if x < cmd.x2 {
		x += 1
	}
	if y > cmd.y2 {
		y -= 1
	} else if y < cmd.y2 {
		y += 1
	}

	return x, y, false
}

func navigate(d *diagram, cmd *command) {
	d.lines[cmd.y1][cmd.x1] += 1
	for x, y, fin := step(cmd.x1, cmd.y1, cmd); !fin; x, y, fin = step(x, y, cmd) {
		d.lines[y][x] += 1
	}
}

func (c *command) isHorizontalOrVertical() bool {
	return c.x1 == c.x2 || c.y1 == c.y2
}

func task1(cmds []*command) (result int) {
	d := NewDiagram(cmds)
	for _, cmd := range cmds {
		if cmd.isHorizontalOrVertical() {
			navigate(d, cmd)
		}
	}
	return d.countOverlaps()
}

func task2(cmds []*command) (result int) {
	d := NewDiagram(cmds)
	for _, cmd := range cmds {
		navigate(d, cmd)
	}
	return d.countOverlaps()
}

func main() {
	input := "input.txt"

	commands := readdata(input)
	fmt.Println("Task 1 - # h&v crossing lines   =  ", task1(commands))
	fmt.Println("Task 2 - # h&v&d crossing lines =  ", task2(commands))
}
