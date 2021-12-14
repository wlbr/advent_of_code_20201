package main

import (
	"fmt"
	"time"
)

type node struct {
	next    *node
	polymer rune
}

func (n *node) walk(visitor func(rune)) {
	for n != nil {
		visitor(n.polymer)
		n = n.next
	}
}

func (n *node) String() string {
	var s string
	n.walk(func(p rune) {
		s += string(p)
	})
	return s
}

func (n *node) splice(another *node) {
	oldnext := n.next
	n.next = another
	another.next = oldnext
}

type rules map[rune]map[rune]rune

func (r rules) String() string {
	s := ""
	for k, v := range r {
		for k2, v2 := range v {
			s += fmt.Sprintf("%c%c -> %c\n", k, k2, v2)
		}
	}
	return s
}

func (r rules) fire(left, right rune) (rune, bool) {
	if precond, ok := r[left]; ok {
		if hit, ok := precond[right]; ok {
			return hit, true
		}
	}
	return 0, false
}

func (r rules) solve(n *node) {
	start := n
	last := start
	n = n.next
	for n != nil {
		if target, ok := r.fire(last.polymer, n.polymer); ok {
			last.splice(&node{polymer: target})
		}
		last = n
		n = n.next
	}
}

const MaxInt = int(^uint(0) >> 1)

func (n *node) score() int {
	singlescores := make(map[rune]int)

	n.walk(func(p rune) {
		singlescores[p]++
	})
	maxv := 0
	minv := MaxInt
	for _, ssc := range singlescores {
		if ssc > maxv {
			maxv = ssc
		}
		if ssc < minv {
			minv = ssc
		}
	}

	return maxv - minv
}

func task1(polymers *node, rules rules) (result int) {
	for i := 1; i <= 10; i++ {
		rules.solve(polymers)
	}

	return polymers.score()
}

func task2(polymers *node, rules rules) (result int) {
	for i := 1; i <= 40; i++ {
		fmt.Println(i)
		rules.solve(polymers)
	}

	return polymers.score()
}

func main() {
	input := "input.txt"

	p, r := readdata(input)
	start := time.Now()
	result := task1(p, r)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(p, r)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
