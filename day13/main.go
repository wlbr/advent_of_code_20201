package main

import (
	"fmt"
	"time"
)

type instruction struct {
	pos      int
	vertical bool
}

type dots struct {
	startx, starty int
	endx, endy     int
	points         [][]bool
}

func (d *dots) get(x, y int) bool {
	return d.points[y+d.starty][x+d.startx]
}

func (d *dots) String() (s string) {
	for y := d.starty; y < d.endy; y++ {
		for x := d.startx; x < d.endx; x++ {
			if d.points[y][x] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += fmt.Sprintln()
	}
	return s
}

func (d *dots) count() (c int) {
	for y := d.starty; y < d.endy; y++ {
		for x := d.startx; x < d.endx; x++ {
			if d.points[y][x] {
				c++
			}
		}
	}
	return c
}

func (d *dots) foldHorizontally(atPos int) {
	for delta := 1; delta < d.endx-d.startx-atPos; delta++ {
		for y := d.starty; y < d.endy; y++ {
			lower := d.points[y][d.startx+atPos+delta]
			upper := d.points[y][d.startx+atPos-delta]
			d.points[y][d.startx+atPos-delta] = lower || upper
		}
	}
	d.endx = d.startx + atPos

}

func (d *dots) foldVertically(atPos int) {
	for delta := 1; delta < d.endy-d.starty-atPos; delta++ {
		for x := d.startx; x < d.endx; x++ {
			lower := d.points[d.starty+atPos+delta][x]
			upper := d.points[d.starty+atPos-delta][x]
			d.points[d.starty+atPos-delta][x] = lower || upper
		}
	}
	d.endy = d.starty + atPos
}

func (d *dots) fold(ops *instruction) {
	if ops.vertical {
		d.foldVertically(ops.pos)
	} else {
		d.foldHorizontally(ops.pos)
	}
}

func task1(dots *dots, ops []*instruction) (result int) {
	dots.fold(ops[0])
	return dots.count()
}

func task2(dots *dots, ops []*instruction) (result int) {
	for _, o := range ops {
		dots.fold(o)
	}
	fmt.Println(dots)
	return dots.count()
}

func main() {
	input := "input.txt"

	dots, ops := readdata(input)

	start := time.Now()
	result := task1(dots, ops)
	fmt.Printf("Task 1 - elapsed Time: %s - #dots after first folds \t = %d \n\n\n", time.Since(start), result)

	//dots, ops = readdata(input)
	start = time.Now()
	result = task2(dots, ops)
	fmt.Printf("Task 2 - elapsed Time: %s - #dots after all folds\t = %d \n", time.Since(start), result)

}
