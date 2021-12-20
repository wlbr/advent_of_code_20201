package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type ImageEnhancerAlgorithm []rune

func (iea ImageEnhancerAlgorithm) String() string {
	s := ""
	for _, r := range iea {
		s += string(r)
	}
	return s
}

func getEncodingForElement(data *InfiniteArray, x, y int) string {
	coords := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	s := ""
	for _, c := range coords {
		s += data.GetString(x+c[0], y+c[1])
	}
	return s
}

func encodingToInt(encoding string) int {
	r := strings.NewReplacer(".", "0", "#", "1")
	b := r.Replace(encoding)
	if d, e := strconv.ParseInt(b, 2, 32); e != nil {
		log.Fatalf("encodingToInt: cannot read binary from '%v'", b)
		return -1
	} else {
		return int(d)
	}
}

func GetValueForElement(data *InfiniteArray, x, y int) int {
	return encodingToInt(getEncodingForElement(data, x, y))
}

func EnhancePixel(iea ImageEnhancerAlgorithm, data *InfiniteArray, x, y int) string {
	return string(iea[GetValueForElement(data, x, y)])
}

func EnhanceImage(algo ImageEnhancerAlgorithm, data *InfiniteArray, times int) (result int) {
	boundaries := 2
	for i := 1; i <= times; i++ {
		var s []string
		for y := -1 - boundaries; y <= len(data.data)+boundaries; y++ {
			line := ""
			for x := -1 - boundaries; x <= len(data.data[0])+boundaries; x++ {
				line += EnhancePixel(algo, data, x, y)
			}
			s = append(s, line)
		}
		defElem := '.'

		if algo[0] == '#' && i%2 != 0 {
			defElem = '#'
		}

		data = NewInfiniteArrayFromString(s, defElem)
	}
	for y := 0; y < data.dimY; y++ {
		for x := 0; x < data.dimX; x++ {
			if data.GetRune(x, y) == '#' {
				result++
			}
		}
	}
	return result
}

func task1(algo ImageEnhancerAlgorithm, data *InfiniteArray) (result int) {
	return EnhanceImage(algo, data, 2)
}

func task2(algo ImageEnhancerAlgorithm, data *InfiniteArray) (result int) {
	return EnhanceImage(algo, data, 50)
}

func main() {
	input := "input.txt"

	algo, data := readdata(input)
	start := time.Now()
	result := task1(algo, data)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	algo, data = readdata(input)
	start = time.Now()
	result = task2(algo, data)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)
}
