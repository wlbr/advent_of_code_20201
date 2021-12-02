package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func navigate1(command string, hpos, depth int) (int, int) {
	segments := strings.Split(command, " ")
	if len(segments) != 2 {
		log.Fatalf("Malformed command '%s'", command)
	}
	lit := segments[0]
	arg, err := strconv.Atoi(segments[1])
	if err != nil {
		log.Fatalf("Unknown argument in command '%s': %v", command, err)
	}

	switch lit {
	case "up":
		depth -= arg
	case "down":
		depth += arg
	case "forward":
		hpos += arg
	default:
		log.Printf("Unknown command '%s'", command)

	}
	return hpos, depth
}

func task1(commands []string) int {
	var hpos, depth int
	for _, command := range commands {
		hpos, depth = navigate1(command, hpos, depth)
	}
	return hpos * depth
}

func navigate2(command string, hpos, depth, aim int) (int, int, int) {
	segments := strings.Split(command, " ")
	if len(segments) != 2 {
		log.Fatalf("Malformed command '%s'", command)
	}
	lit := segments[0]
	arg, err := strconv.Atoi(segments[1])
	if err != nil {
		log.Fatalf("Unknown argument in command '%s': %v", command, err)
	}

	switch lit {
	case "up":
		aim -= arg
	case "down":
		aim += arg
	case "forward":
		hpos += arg
		depth += (aim * arg)
	default:
		log.Printf("Unknown command '%s'", command)

	}
	return hpos, depth, aim
}

func task2(commands []string) int {
	var hpos, depth, aim int
	for _, command := range commands {
		hpos, depth, aim = navigate2(command, hpos, depth, aim)
	}
	log.Printf("hpos: %d, depth: %d, aim: %d", hpos, depth, aim)
	return hpos * depth
}

func main() {
	input := "input.txt"

	ids := readdata(input)
	fmt.Println("Task 1 - # increasing depths   =  ", task1(ids))
	fmt.Println("Task 1 - # sliding inc. depths =  ", task2(ids))
}
