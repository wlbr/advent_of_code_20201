package main

import (
	"fmt"
	"log"
)

func navigate(command string, hpos, depth, aim int) (int, int, int) {
	lit, arg := splitCommand(command)

	switch lit {
	case "up":
		aim -= arg
	case "down":
		aim += arg
	case "forward":
		hpos += arg
		depth += (arg * aim)
	default:
		log.Printf("Unknown command '%s'", command)
	}
	return hpos, depth, aim
}

func task1(commands []string) int {
	var hpos, aim int
	for _, command := range commands {
		hpos, _, aim = navigate(command, hpos, 1, aim)
	}
	return hpos * aim
}

func task2(commands []string) int {
	var hpos, depth, aim int
	for _, command := range commands {
		hpos, depth, aim = navigate(command, hpos, depth, aim)
	}
	return hpos * depth
}

func main() {
	input := "input.txt"

	ids := readdata(input)
	fmt.Println("Task 1 - # increasing depths   =  ", task1(ids))
	fmt.Println("Task 1 - # sliding inc. depths =  ", task2(ids))
}
