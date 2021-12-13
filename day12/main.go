package main

import (
	"fmt"
	"strings"
	"time"
)

type node struct {
	name        string
	isLarge     bool
	connections []*node
}

func newNode(name string) *node {
	return &node{name: name, isLarge: strings.ToUpper(name) == name, connections: []*node{}}
}

func (n *node) String() string {
	return fmt.Sprintf("%s: %v", n.name, n.connections)
}

func (n *node) addConnection(c *node) {
	n.connections = append(n.connections, c)
	c.connections = append(c.connections, n)
}

type plan struct {
	start, end *node
	node       map[string]*node
	allNodes   []*node
}

func newPlan(start, end *node) *plan {
	return &plan{node: make(map[string]*node)}
}

func printAllPaths(paths [][]*node) {
	for _, p := range paths {
		for i, sp := range p {
			fmt.Printf("%s", sp.name)
			if i < len(p)-1 {
				fmt.Print(",")
			} else {
				fmt.Println()
			}
		}
	}
}

func (p *plan) add(name string) *node {
	if n, ok := p.node[name]; ok {
		return n
	}
	n := newNode(name)
	p.node[name] = n
	p.allNodes = append(p.allNodes, n)
	return n
}

func (p *plan) get(name string) *node {
	return p.node[name]
}

func count(n *node, path []*node) int {
	c := 0
	for _, pn := range path {
		if n == pn {
			c++
		}
	}
	return c
}

type visitChecker func(n *node, currentpath []*node, anyoneDoubleVisited bool) (bool, bool)

func canVisit1(n *node, currentpath []*node, anyoneDoubleVisited bool) (bool, bool) {
	if count(n, currentpath) == 0 || n.isLarge {
		return true, false
	}
	return false, false
}

func canVisit2(n *node, currentpath []*node, anyoneDoubleVisited bool) (bool, bool) {
	count := count(n, currentpath)
	if count == 0 || n.isLarge {
		return true, anyoneDoubleVisited
	}
	if count == 1 && !anyoneDoubleVisited && n.name != "start" && n.name != "end" {
		return true, true
	}
	return false, anyoneDoubleVisited
}

func (p *plan) travel(paths [][]*node, currentpath []*node, startnode *node, canvisit visitChecker, anyTwiceVisited bool) [][]*node {
	currentpath = append(currentpath, startnode)
	if startnode.name == p.end.name {
		paths = append(paths, currentpath)
		return paths
	}
	for _, c := range startnode.connections {
		cvisit, twiceAlready := canvisit(c, currentpath, anyTwiceVisited)
		if cvisit {
			paths = p.travel(paths, currentpath, c, canvisit, twiceAlready)
		}
	}
	return paths
}

func task1(plan *plan) (result int) {
	var paths [][]*node
	paths = plan.travel(paths, nil, plan.start, canVisit1, false)

	return len(paths)
}

func task2(plan *plan) (result int) {
	var paths [][]*node
	paths = plan.travel(paths, nil, plan.start, canVisit2, false)

	return len(paths)
}

func task3(plan *plan) (result int) {

	results := make(chan int, len(plan.start.connections))

	for _, c := range plan.start.connections {
		go func(res chan int, start *node) {
			res <- len(plan.travel(nil, []*node{plan.start}, start, canVisit2, false))
		}(results, c)
	}

	for i := 0; i < len(plan.start.connections); i++ {
		result += <-results
	}

	close(results)
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)

	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - # Paths\t\t\t= %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)

	fmt.Printf("Task 2 - elapsed Time: %s - # Paths\t\t\t= %d \n", time.Since(start), result)

	start = time.Now()
	result = task3(data)

	fmt.Printf("Task 3 - elapsed Time: %s - # Paths in parallel\t= %d \n", time.Since(start), result)

}
