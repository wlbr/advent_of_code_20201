package main

import (
	"fmt"
)

func task1(depths []int) (increasingdepths int) {
	last := -1
	for _, d := range depths {
		if last != -1 && d > last {
			increasingdepths++
		}
		last = d
	}
	return increasingdepths
}

func task2(depths []int) (increasingdepths int) {
	lastsum := depths[2] + depths[1] + depths[0]

	for i := 3; i < len(depths); i++ {
		current := depths[i] + depths[i-1] + depths[i-2]
		if current > lastsum {
			increasingdepths++
			lastsum = current
		}
	}
	return increasingdepths
}

func task2_dumbo(depths []int) (increasingdepths int) {
	var (
		sum1, sum2, sum3                                  int = -1, -1, -1
		slidingcounter1, slidingcounter2, slidingcounter3 int = 0, -1, -2
		last1, last2, last3                               int = -1, -1, -1
	)

	for _, d := range depths {
		slidingcounter1++
		slidingcounter2++
		slidingcounter3++

		sum1 = sum1 + d
		sum2 = sum2 + d
		sum3 = sum3 + d

		if slidingcounter1%3 == 0 {
			if slidingcounter1 > 0 {
				last1 = sum1
				if sum1 > last3 && last3 >= 0 {
					increasingdepths++
				}
				sum1 = 0
			}
		}

		if slidingcounter2%3 == 0 {
			if slidingcounter2 > 0 {
				last2 = sum2
				if sum2 > last1 && last1 >= 0 {
					increasingdepths++
				}
			}
			sum2 = 0
		}

		if slidingcounter3%3 == 0 {
			if slidingcounter3 > 0 {
				last3 = sum3
				if sum3 > last2 && last2 >= 0 {
					increasingdepths++
				}
			}
			sum3 = 0
		}
	}
	return increasingdepths
}

func main() {
	input := "input.txt"

	ids := readdata(input)
	fmt.Println("Task 1 - # increasing depths   =  ", task1(ids))
	fmt.Println("Task 1 - # sliding inc. depths =  ", task2(ids))
}
