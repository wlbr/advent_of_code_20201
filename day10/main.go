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
	var corruptingrunes []rune

	for _, l := range alllines {
		v, r, _ := isValid(l)
		if v == 1 {
			corruptingrunes = append(corruptingrunes, r)
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

func complete(in string) (out string) {
	for _, c := range in {
		if c == '(' {
			out = ")" + out
		} else if c == '[' {
			out = "]" + out
		} else if c == '{' {
			out = "}" + out
		} else if c == '<' {
			out = ">" + out
		}
	}
	return out
}

func task2(alllines []string) int {
	var missingopenings []string
	for _, l := range alllines {
		v, _, m := isValid(l)
		if v == -1 {
			missingopenings = append(missingopenings, m.content)
		}
	}

	missingclosings := []string{}
	for _, mops := range missingopenings {
		missingclosings = append(missingclosings, complete(mops))
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
