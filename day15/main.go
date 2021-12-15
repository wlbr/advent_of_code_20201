package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beefsack/go-astar"
)

type Risklevel struct {
	x, y  int
	value int
	cave  *Cave
}
type Cave struct {
	levels  [][]*Risklevel
	dimx    int
	dimy    int
	maxrisk int
}

func (c *Cave) String() string {
	var result string
	for _, row := range c.levels {
		for _, cell := range row {
			result += fmt.Sprintf("%d", cell.value)
		}
		result += "\n"
	}
	return result
}

func (c *Cave) enlargeCave(factor int) *Cave {
	nc := &Cave{}
	nc.dimx = c.dimx * factor
	nc.dimy = c.dimy * factor
	nc.levels = make([][]*Risklevel, nc.dimy)
	for yf := 0; yf < factor; yf++ {
		offsety := c.dimy * yf
		for y := 0; y < c.dimy; y++ {
			nc.levels[y+offsety] = make([]*Risklevel, nc.dimx)
			for xf := 0; xf < factor; xf++ {
				offsetx := c.dimx * xf
				for x := 0; x < c.dimx; x++ {
					v := ((c.levels[y][x].value + xf + yf - 1) % 9) + 1
					nc.levels[y+offsety][x+offsetx] = &Risklevel{x + offsetx, y + offsety, v, nc}
				}
			}
		}
	}
	return nc
}

func (r *Risklevel) getAdjacentValues() (levels []astar.Pather) {
	x := r.x
	y := r.y
	dimx := r.cave.dimx
	dimy := r.cave.dimy
	if x > 0 {
		levels = append(levels, r.cave.levels[y][x-1])
	}
	if x < dimx-1 {
		levels = append(levels, r.cave.levels[y][x+1])
	}
	if y > 0 {
		levels = append(levels, r.cave.levels[y-1][x])
	}
	if y < dimy-1 {
		levels = append(levels, r.cave.levels[y+1][x])
	}
	return levels
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (r *Risklevel) ManhattanDistance(to *Risklevel) float64 {
	return float64(abs(r.x-to.x) + abs(r.y-to.y))
}

func (r *Risklevel) PathNeighbors() []astar.Pather {
	return r.getAdjacentValues()
}

func (t *Risklevel) PathNeighborCost(to astar.Pather) float64 {
	tor := to.(*Risklevel)
	return float64(tor.value)
}

func (t *Risklevel) PathEstimatedCost(to astar.Pather) float64 {
	return t.ManhattanDistance(to.(*Risklevel))
}

func task1(cave *Cave) (result int) {
	_, distance, found := astar.Path(cave.levels[0][0], cave.levels[cave.dimy-1][cave.dimx-1])

	if !found {
		log.Println("Could not find path")
	}

	return int(distance)
}

func task2(cave *Cave) (result int) {
	c := cave.enlargeCave(5)

	_, distance, found := astar.Path(c.levels[0][0], c.levels[c.dimy-1][c.dimx-1])

	if !found {
		log.Println("Could not find path")
	}

	return int(distance)
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - Path costs \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s - Path costs \t = %d \n", time.Since(start), result)
}
