package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

func isValid(line string) (int, lastsign rune, mem *stack) {
	opening := "([{<"
	closing := ")]}>"

	mem = &stack{}
	for _, c := range line {
		if strings.Contains(opening, string(c)) {
			mem.push(c)
		} else if i := strings.Index(closing, string(c)); i >= 0 {

			sp := mem.pop()
			if sp != opening[i] {
				return 1, c, mem
			}
		}
	}
	if !mem.empty() {
		return -1, rune(mem.peek()), mem
	}
	return 0, ' ', mem
}

func score(corruptingrunes []rune) int {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	s := 0
	for _, r := range corruptingrunes {
		if p, ok := points[r]; ok {
			s += p
		} else {
			log.Fatal("Invalid closing rune")
		}
	}
	return s
}

func task1(alllines []string) (result int) {
	var corruptlines, validlines, incompletelines []string
	var corruptids, validids, incompleteids []int
	var corruptingrunes []rune

	for p, l := range alllines {
		v, r, _ := isValid(l)
		switch v {
		case 0:
			validlines = append(validlines, l)
			validids = append(validids, p)
		case 1:
			corruptlines = append(corruptlines, l)
			corruptids = append(corruptids, p)
			corruptingrunes = append(corruptingrunes, r)
		case -1:
			incompletelines = append(incompletelines, l)
			incompleteids = append(incompleteids, p)
		}
	}

	return score(corruptingrunes)
}

func score2(missinglineclosings []string) int {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var scores []int
	for _, m := range missinglineclosings {
		s := 0
		for _, c := range m {
			if p, ok := points[c]; ok {
				s = s*5 + p
			} else {
				log.Fatalf("Invalid closing rune: %s", string(c))
			}
		}
		scores = append(scores, s)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func task2(alllines []string) int {
	var corruptlines, validlines, incompletelines []string
	var corruptids, validids, incompleteids []int

	var missingopenings []string
	for p, l := range alllines {
		v, _, m := isValid(l)
		switch v {
		case 0:
			validlines = append(validlines, l)
			validids = append(validids, p)
		case 1:
			corruptlines = append(corruptlines, l)
			corruptids = append(corruptids, p)
		case -1:
			incompletelines = append(incompletelines, l)
			incompleteids = append(incompleteids, p)
			missingopenings = append(missingopenings, m.content)
		}
	}
	reverse := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	missingclosings := []string{}
	for _, mops := range missingopenings {
		rops := ""
		for _, m := range mops {
			rops = string(reverse[m]) + rops
		}
		missingclosings = append(missingclosings, rops)
	}
	return score2(missingclosings)
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	fmt.Printf("Task 1 - elapsed Time: %s - Syntax error score       = %d \n", time.Since(start), task1(data))

	start = time.Now()
	fmt.Printf("Task 2 - elapsed Time: %s - Mid of completion scores = %d \n", time.Since(start), task2(data))
}
