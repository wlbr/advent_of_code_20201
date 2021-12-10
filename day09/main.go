package main

import (
	"fmt"
	"sort"
	"time"
)

func Min(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

func riskLevel(i int) int {
	return 1 + i
}

func getAdjacentValues(heights [][]int, x, y int) (adjacents []int) {

	dimx := len(heights[0])
	dimy := len(heights)
	if x > 0 {
		adjacents = append(adjacents, heights[y][x-1])
	}
	if x < dimx-1 {
		adjacents = append(adjacents, heights[y][x+1])
	}
	if y > 0 {
		adjacents = append(adjacents, heights[y-1][x])
	}
	if y < dimy-1 {
		adjacents = append(adjacents, heights[y+1][x])
	}
	return adjacents
}

func task1(heights [][]int) (sum int) {
	var lowpoints []int

	for y, row := range heights {
		for x, height := range row {
			adjacentsHeights := getAdjacentValues(heights, x, y)
			if len(adjacentsHeights) > 0 {
				if Min(adjacentsHeights...) > height {
					lowpoints = append(lowpoints, height)
				}
			}
		}
	}
	for _, v := range lowpoints {
		sum += riskLevel(v)
	}

	return sum
}

type pos struct {
	x, y int
}

func (p pos) equals(o pos) bool {
	return p.x == o.x && p.y == o.y
}

func (p pos) getAdjacentPositions(dimx, dimy int) (adjacents []pos) {
	if p.x > 0 {
		adjacents = append(adjacents, pos{p.x - 1, p.y})
	}
	if p.x < dimx-1 {
		adjacents = append(adjacents, pos{p.x + 1, p.y})
	}
	if p.y > 0 {
		adjacents = append(adjacents, pos{p.x, p.y - 1})
	}
	if p.y < dimy-1 {
		adjacents = append(adjacents, pos{p.x, p.y + 1})
	}
	return adjacents
}

type basin struct {
	spots []pos
}

func (b *basin) contains(p pos) bool {
	for _, v := range b.spots {
		if v.equals(p) {
			return true
		}
	}
	return false
}

func (b *basin) size() int {
	return len(b.spots)
}

func (b *basin) add(p pos) {
	b.spots = append(b.spots, p)
}

func (b *basin) merge(other *basin) {
	if b != other {
		b.spots = append(b.spots, other.spots...)
		other.spots = nil
	}
}

func (p pos) getAdjacentBasins(basins []*basin, dimx, dimy int) map[*basin]*struct{} {
	adjacentPositions := p.getAdjacentPositions(dimx, dimy)
	adjacentBasins := make(map[*basin]*struct{})
	for _, ap := range adjacentPositions {
		for _, b := range basins {
			if b.contains(ap) {
				adjacentBasins[b] = nil
			}
		}
	}
	return adjacentBasins
}

func findBasins(heights [][]int) (basins []*basin) {
	dimx := len(heights[0])
	dimy := len(heights)

	for y, row := range heights {
		for x, height := range row {
			if height != 9 {
				p := pos{x, y}
				adjacentBasins := p.getAdjacentBasins(basins, dimx, dimy)

				switch len(adjacentBasins) {
				case 0:
					basins = append(basins, &basin{spots: []pos{p}})
				case 1:
					for b := range adjacentBasins {
						b.add(p)
					}
				default:
					var first *basin
					for b := range adjacentBasins {
						if nil == first {
							first = b
						} else {
							first.merge(b)
						}
					}
					first.add(p)
				}
			}
		}
	}

	var unemtybasins []*basin
	for _, b := range basins {
		if b.size() > 0 {
			unemtybasins = append(unemtybasins, b)
		}
	}
	return unemtybasins
}

func task2(heights [][]int) (result int) {
	basins := findBasins(heights)
	var basinsSizes []int

	for _, b := range basins {
		basinsSizes = append(basinsSizes, b.size())
	}

	if len(basinsSizes) >= 3 {
		sort.Ints(basinsSizes)
		result = basinsSizes[len(basinsSizes)-1] * basinsSizes[len(basinsSizes)-2] * basinsSizes[len(basinsSizes)-3]

	}
	return result
}

func main() {
	input := "input.txt"

	heights := readdata(input)
	start := time.Now()
	result := task1(heights)
	fmt.Printf("Task 1 - elapsed Time: %s - Sum of risk levels     = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(heights)
	fmt.Printf("Task 2 - elapsed Time: %s - Products of basinsizes = %d \n", time.Since(start), result)
}
