package main

import (
	"fmt"
)

type Fish struct {
	counter int
}

func NewFish(count int) *Fish {
	return &Fish{counter: count}
}

type Basket struct {
	age    int
	number int
}

func countFish(baskets []*Basket) int {
	var count int
	for _, b := range baskets {
		count += b.number
	}
	return count
}

func age(baskets []*Basket) []*Basket {
	var newBaskets []*Basket
	for a := 0; a < len(baskets); a++ {
		newBaskets = append(newBaskets, &Basket{age: a, number: 0})
	}
	for i := len(baskets) - 2; i >= 0; i-- {
		newBaskets[i].number = baskets[i+1].number
		if i == 0 {
			newBaskets[len(baskets)-1].number = baskets[0].number
			newBaskets[6].number += baskets[0].number
		}
	}
	return newBaskets
}

func task1and2(fishes []*Fish, days int) (result int) {
	baskets := make([]*Basket, 9)
	for a := 0; a < len(baskets); a++ {
		baskets[a] = &Basket{age: a, number: 0}
	}
	for _, f := range fishes {
		baskets[f.counter].number++
	}
	for day := 0; day < days; day++ {
		baskets = age(baskets)
	}
	return countFish(baskets)
}

func main() {
	input := "input.txt"

	fishes := readdata(input)
	days := 80
	fmt.Printf("Task 1 - # fishes after %d days   = %d \n", days, task1and2(fishes, days))
	days = 256
	fmt.Printf("Task 2 - # fishes after %d days  = %d \n", days, task1and2(fishes, days))
}
