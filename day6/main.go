package main

import (
	"fmt"
)

type basket struct {
	age    int
	number int
}

func countFish(baskets []*basket) int {
	var count int
	for _, b := range baskets {
		count += b.number
	}
	return count
}

func age(baskets []*basket) []*basket {
	var NewBaskets []*basket
	for a := 0; a < len(baskets); a++ {
		NewBaskets = append(NewBaskets, &basket{age: a, number: 0})
	}
	for i := len(baskets) - 2; i >= 0; i-- {
		NewBaskets[i].number = baskets[i+1].number
		if i == 0 {
			NewBaskets[len(baskets)-1].number = baskets[0].number
			NewBaskets[6].number += baskets[0].number
		}
	}
	return NewBaskets
}

func task1and2(fishes []int, days int) (result int) {
	baskets := make([]*basket, 9)
	for a := 0; a < len(baskets); a++ {
		baskets[a] = &basket{age: a, number: 0}
	}
	for _, f := range fishes {
		baskets[f].number++
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
