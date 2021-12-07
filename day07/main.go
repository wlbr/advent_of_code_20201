package main

import (
	"fmt"
	"math"
)

const MaxInt = int(^uint(0) >> 1)

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

func Move(current, target int, progressive bool) int {
	distance := int(math.Abs(float64(current - target)))
	if !progressive { // task1
		return distance
	}
	c := 0 // task2
	for i := 1; i <= distance; i++ {
		c += i
	}
	return c
}

func Costs(crabs []int, target int, progressive bool) int {
	count := 0

	for _, current := range crabs {
		count += Move(current, target, progressive)
	}
	return count
}

func Cheapest(crabs []int, progressive bool) (target, cost int) {
	min, max := MinMax(crabs...)
	fmt.Printf("%v - Min %d Max %d \n", crabs, min, max)
	cost = MaxInt

	for i := min; i < max; i++ {
		c := Costs(crabs, i, progressive)
		if c < cost {
			target = i
			cost = c
		}
	}
	return target, cost
}

func task1(crabs []int) (target, cost int) {
	return Cheapest(crabs, false)
}

func task2(crabs []int) (target, cost int) {
	return Cheapest(crabs, true)
}

func main() {
	input := "input.txt"

	crabpositions := readdata(input)

	target, cost := task1(crabpositions)
	fmt.Printf("Task 1 - # cheapest position is %d costing %d fuel \n", target, cost)
	target, cost = task2(crabpositions)
	fmt.Printf("Task 2 - # cheapest position is %d costing %d fuel \n", target, cost)

}
