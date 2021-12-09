package main

import (
	"fmt"
	"sort"
)

func checksimple(id string) (value int) {
	switch len(id) {
	case 2:
		value = 1
	case 3:
		value = 7
	case 4:
		value = 4
	case 7:
		value = 8
	default:
		value = -1
	}
	return value
}

func task1(signals, patterns [][]string) (result int) {
	for _, line := range patterns {
		for _, p := range line {
			if v := checksimple(p); v != -1 {
				result++
			}
		}
	}
	return result
}

func checkSixLengths(encode Encoder, sixornineorzero []string) {
	for _, n := range sixornineorzero {
		if len(and(encode[4], n)) == 4 {
			encode[9] = n
		} else if len(and(encode[7], n)) == 3 {
			encode[0] = n
		} else {
			if len(and(encode[7], n)) == 2 {
				encode[6] = n
			}
		}
	}
}

func checkFiveLengths(encode Encoder, twoorthreeorfive []string) {
	for _, n := range twoorthreeorfive {
		if len(xor(encode[1], n)) == 3 {
			encode[3] = n
		} else if len(xor(encode[4], n)) == 3 {
			encode[5] = n
		} else {
			if len(xor(encode[4], n)) == 5 {
				encode[2] = n
			}
		}
	}
}

func Deduction(signals []string) (Decoder, Encoder) {
	decode := make(Decoder)
	encode := make(Encoder)
	lenghts := make(map[int][]string)

	sort.Slice(signals, func(i, j int) bool {
		return len(signals[i]) < len(signals[j])
	})

	for _, p := range signals {
		if v := checksimple(p); v != -1 {
			encode[v] = p
		}
		lenghts[len(p)] = append(lenghts[len(p)], p)
	}

	checkSixLengths(encode, lenghts[6])
	checkFiveLengths(encode, lenghts[5])

	for k, v := range encode {
		decode[v] = k
	}
	return decode, encode
}

type Decoder map[string]int

func (c Decoder) String() string {
	r := ""
	for p, v := range c {
		r += fmt.Sprintf("Found %s = %d\n", p, v)
	}
	return r
}

type Encoder map[int]string

func (c Encoder) String() string {
	r := ""
	for p, v := range c {
		r += fmt.Sprintf("Found %d = %s\n", p, v)
	}
	return r
}

func task2(signals, patterns [][]string) (result int) {
	sum := 0
	for i := 0; i < len(signals); i++ {
		decode, _ := Deduction(signals[i])
		r := 1000*decode[patterns[i][0]] +
			100*decode[patterns[i][1]] +
			10*decode[patterns[i][2]] +
			decode[patterns[i][3]]
		sum += r
	}

	return sum
}

func main() {
	input := "input.txt"
	signals, patterns := readdata(input)
	fmt.Println("Task 1 - # sum of count 1,4,7,8   =  ", task1(signals, patterns))
	fmt.Println("Task 1 - # decoded sum            =  ", task2(signals, patterns))
}
