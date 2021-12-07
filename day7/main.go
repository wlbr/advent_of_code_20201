package main

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Min(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

func MinMax(v ...int) (min, max int) {
	max = v[0]
	min = v[0]
	for _, e := range v {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	return min, max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Move(current, target int) int {
	return Abs(current - target)
}

func Costs(positions []int, target int) int {
	count := 0

	for _, current := range positions {
		count += Move(current, target)
	}
	return count
}

func Cheapest(positions []int) (target, cost int) {
	min, max := MinMax(positions...)
	fmt.Printf("%v - Min %d Max %d \n", positions, min, max)
	cost = MaxInt

	for i := min; i < max; i++ {
		c := Costs(positions, i)
		if c < cost {
			target = i
			cost = c
		}
	}

	return target, cost
}

func ProgressiveMove(current, target int) int {
	distance := Abs(current - target)
	c := 0

	for i := 1; i <= distance; i++ {
		c += i
	}
	return c
}

func ProgressiveCosts(positions []int, target int) int {
	count := 0

	for _, current := range positions {
		count += ProgressiveMove(current, target)
	}
	return count
}

func Cheapest2(positions []int) (target, cost int) {
	min, max := MinMax(positions...)
	fmt.Printf("%v - Min %d Max %d \n", positions, min, max)
	cost = MaxInt

	for i := min; i < max; i++ {
		c := ProgressiveCosts(positions, i)
		fmt.Printf("target: %d   costs: %d\n", i, c)
		if c < cost {
			target = i
			cost = c
		}
	}
	return target, cost
}

func task1(positions []int) (target, cost int) {
	return Cheapest(positions)
}

func task2(positions []int) (target, cost int) {
	return Cheapest2(positions)
}

func main() {
	input := "input.txt"

	fishes := readdata(input)

	target, cost := task1(fishes)
	fmt.Printf("Task 1 - # cheapest position is %d costing %d fuel \n", target, cost)
	target, cost = task2(fishes)
	fmt.Printf("Task 1 - # cheapest position is %d costing %d fuel \n", target, cost)

}
