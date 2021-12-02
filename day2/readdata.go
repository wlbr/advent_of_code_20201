package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) (commands []string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return commands
}
