package main

import (
	"fmt"
	"math/big"
)

type basket struct {
	age    int
	number *big.Int
}

func countFish(baskets []*basket) *big.Int {
	count := big.NewInt(0)

	for _, b := range baskets {
		count = count.Add(count, b.number)
	}
	return count
}

func age(baskets []*basket) []*basket {
	NewBaskets := append(baskets[1:], baskets[0])
	NewBaskets[6].number = NewBaskets[6].number.Add(NewBaskets[6].number, NewBaskets[len(NewBaskets)-1].number)
	return NewBaskets
}

func task1and2(fishes []int, days int) (result *big.Int) {
	baskets := make([]*basket, 9)
	for a := 0; a < len(baskets); a++ {
		baskets[a] = &basket{age: a, number: big.NewInt(0)}
	}
	for _, f := range fishes {
		baskets[f].number = baskets[f].number.Add(baskets[f].number, big.NewInt(1))
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
