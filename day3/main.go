package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func mostCommon(binaries []string, pos int) string {
	var (
		count0, count1 int = 0, 0
	)
	for _, w := range binaries {
		switch w[pos] {
		case '0':
			count0++
		default:
			count1++
		}
	}
	if count1 >= count0 {
		return "1"
	} else {
		return "0"
	}
}

func leastCommon(binaries []string, pos int) string {
	if mostCommon(binaries, pos) == "1" {
		return "0"
	} else {
		return "1"
	}
}

func filter(binaries []string, bit int, criteria byte) (fitting []string) {
	for _, num := range binaries {
		if num[bit] == criteria {
			fitting = append(fitting, num)
		}
	}
	return fitting
}

func gamma(bin []string) int {
	commons := ""
	wordlength := len(bin[0])
	for i := 0; i < wordlength; i++ {
		commons += mostCommon(bin, i)
	}
	gamma, _ := strconv.ParseInt(commons, 2, 64)
	return int(gamma)
}

func epsilon(bin []string) int {
	xormask := int(math.Pow(float64(2), float64(len(bin[0])))) - 1
	return xormask ^ gamma(bin)
}

func oxygengenerator(binaries []string) (result int) {
	set := binaries
	for i := 0; i < len(binaries[0]); i++ {
		criteria := mostCommon(set, i)
		set = filter(set, i, criteria[0])
		if len(set) == 1 {
			oxy, _ := strconv.ParseInt(set[0], 2, 64)
			return int(oxy)
		}
	}
	log.Fatal("oxygen: no solution found.")
	return 0
}

func co2scrubber(binaries []string) (result int) {
	set := binaries
	for i := 0; i < len(binaries[0]); i++ {
		criteria := leastCommon(set, i)
		set = filter(set, i, criteria[0])
		if len(set) == 1 {
			oxy, _ := strconv.ParseInt(set[0], 2, 64)
			return int(oxy)
		}
	}
	log.Fatal("co2: no solution found.")
	return 0
}

func task1(binaries []string) (result int) {
	g := gamma(binaries)
	e := epsilon(binaries)
	return g * e
}

func task2(binaries []string) (result int) {
	oxy := oxygengenerator(binaries)
	co2 := co2scrubber(binaries)
	return oxy * co2
}

func main() {
	input := "input.txt"

	ids := readdata(input)
	fmt.Println("Task 1 - # increasing depths   =  ", task1(ids))
	fmt.Println("Task 1 - # sliding inc. depths =  ", task2(ids))
}
