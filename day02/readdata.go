package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func splitCommand(command string) (string, int) {
	segments := strings.Split(command, " ")
	if len(segments) != 2 {
		log.Fatalf("Malformed command '%s'", command)
	}
	lit := segments[0]
	arg, err := strconv.Atoi(segments[1])
	if err != nil {
		log.Fatalf("Unknown argument in command '%s': %v", command, err)
	}
	return lit, arg
}
