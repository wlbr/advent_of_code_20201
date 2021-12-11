package main

import (
	"fmt"
	"time"
)

type alloctos [][]*octopus

func (o alloctos) String() string {
	s := ""
	for _, l := range o {
		s += fmt.Sprintf("%v\n", l)
	}
	return s
}

type octopus struct {
	level   int
	flashed bool
	x, y    int
}

func (o octopus) String() string {
	return fmt.Sprintf("%d", o.level)
}

func (o *alloctos) flash(x, y int) (result int) {
	return 0
}

func (allo alloctos) getAdjacentPositions(x, y int) (adjacents [][]int) {
	dimy := len(allo)
	dimx := len(allo[0])

	if x > 0 && y > 0 {
		adjacents = append(adjacents, []int{x - 1, y - 1})
	}
	if y > 0 {
		adjacents = append(adjacents, []int{x, y - 1})
	}
	if x < dimx-1 && y > 0 {
		adjacents = append(adjacents, []int{x + 1, y - 1})
	}
	if x > 0 {
		adjacents = append(adjacents, []int{x - 1, y})
	}
	if x < dimx-1 {
		adjacents = append(adjacents, []int{x + 1, y})
	}
	if x > 0 && y < dimy-1 {
		adjacents = append(adjacents, []int{x - 1, y + 1})
	}
	if y < dimy-1 {
		adjacents = append(adjacents, []int{x, y + 1})
	}
	if x < dimx-1 && y < dimy-1 {
		adjacents = append(adjacents, []int{x + 1, y + 1})
	}
	return adjacents
}

func (allo alloctos) getAdjacentOctopusses(x, y int) (adjacents []*octopus) {
	pos := allo.getAdjacentPositions(x, y)
	for _, p := range pos {
		adjacents = append(adjacents, allo[p[1]][p[0]])
	}

	return adjacents
}

func (alloct alloctos) incAll() {
	for y := 0; y < len(alloct); y++ {
		for x := 0; x < len(alloct[y]); x++ {
			alloct[y][x].level++
		}
	}
}

func (alloct alloctos) resetFlashed() {
	for _, row := range alloct {
		for _, o := range row {
			if o.flashed {
				o.reset()
			}
		}
	}
}

func (o *octopus) checkAndFlash(alloct alloctos) (changed bool) {
	if !o.flashed && o.level > 9 {
		o.flashed = true
		flashes++
		adjaoctos := alloct.getAdjacentOctopusses(o.x, o.y)
		for x := 0; x < len(adjaoctos); x++ {
			adjaoctos[x].level++
		}

		for _, ao := range adjaoctos {
			changed = changed || ao.checkAndFlash(alloct)
		}
		return changed
	}

	return changed
}

func (o *octopus) reset() {
	o.flashed = false
	o.level = 0
}

func (alloct alloctos) allFlashed() bool {
	for _, row := range alloct {
		for _, o := range row {
			if !o.flashed {
				return false
			}
		}
	}
	return true
}

func (alloct alloctos) doStep() (allflashed bool) {
	alloct.incAll()
	changed := true
	for changed {
		changed = false
		for _, row := range alloct {
			for _, o := range row {
				changed = changed || o.checkAndFlash(alloct)
			}
		}
		if alloct.allFlashed() {
			return true
		}
		alloct.resetFlashed()
	}
	return false
}

var flashes int

func task1(alloct alloctos) (result int) {
	for i := 1; i <= 100; i++ {
		alloct.doStep()
	}
	return flashes
}

func task2(alloct alloctos) (result int) {
	for i := 1; i <= 200000; i++ {
		if alloct.doStep() {
			return i
		}
	}
	return -1
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - flashes after 100 steps     \t = %d \n", time.Since(start), result)

	data = readdata(input)
	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s - number of steps till all flash \t = %d \n", time.Since(start), result)

}
