package main

import (
	"fmt"
	"math"
	"time"
)

type rules map[string]string

func (r rules) String() string {
	s := ""
	for k, v := range r {
		s += fmt.Sprintf("%s -> %s\n", k, v)
	}
	return s
}

func (r rules) fire(precond string) (string, bool) {
	if hit, ok := r[precond]; ok {
		return hit, true
	}
	return "", false
}

func (r rules) solveIteratively(nn string) string {
	n := nn
	for i := 1; i < len(n); i++ {
		if target, ok := r.fire(n[i-1 : i+1]); ok {
			n = n[:i] + target + n[i:]
			i += len(target)
		}
	}
	return n
}

type polymer struct {
	name  string
	depth int
}

type polycounts struct {
	counts map[rune]int64
}

var cache map[polymer]polycounts

func addcaches(a, b map[rune]int64) map[rune]int64 {
	res := make(map[rune]int64)
	for k, v := range a {
		res[k] += v
	}
	for k, v := range b {
		res[k] += v
	}
	return res
}

func gen(rules map[string]string, cache map[polymer]polycounts, in string, dep int) map[rune]int64 {

	if dep == 40 {
		return nil
	}
	if v, ok := cache[polymer{in, dep}]; ok {
		return v.counts
	}
	if r, ok := rules[in]; ok {
		left := gen(rules, cache, string(in[0])+r, dep+1)
		right := gen(rules, cache, r+string(in[1]), dep+1)
		counts := addcaches(left, right)
		counts[rune(r[0])]++
		cache[polymer{in, dep}] = polycounts{counts: counts}
		return counts
	}
	return nil
}

const MaxInt = int(^uint(0) >> 1)

func stringScore(polymers string) int {
	singlescores := make(map[string]int)

	for _, c := range polymers {
		singlescores[string(c)]++
	}

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

func solveRecursively(polymers string, rules rules, dep int) int {
	cache = make(map[polymer]polycounts)

	c := make(map[rune]int64)
	for _, v := range polymers {
		c[v]++
	}
	for i := 0; i < len(polymers)-1; i++ {
		s := gen(rules, cache, polymers[i:i+2], 0)
		c = addcaches(c, s)
	}
	var mx, mi int64 = 0, math.MaxInt64
	for _, v := range c {
		if v > mx {
			mx = v
		}
		if v < mi {
			mi = v
		}
	}
	return int(mx - mi)
}

func task1(polymers string, rules rules) (result int) {
	poly := polymers
	for i := 1; i <= 10; i++ {
		poly = rules.solveIteratively(poly)
	}
	return stringScore(poly)
}

func task2(polymers string, rules rules) (result int) {
	return solveRecursively(polymers, rules, 10)
}

func main() {
	input := "input.txt"

	p, r := readdata(input)

	start := time.Now()
	result := task1(p, r)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	p, r = readdata(input)
	start2 := time.Now()
	result2 := task2(p, r)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start2), result2)

}
