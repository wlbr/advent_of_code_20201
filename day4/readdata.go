package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func (b *bingo) readchoices(scanner *bufio.Scanner) {
	scanner.Scan()
	choicesline := scanner.Text()
	choicesstrings := strings.Split(choicesline, ",")
	for _, cs := range choicesstrings {
		c, err := strconv.Atoi(cs)
		if err != nil {
			log.Fatalf("Error converting choice '%s' to int:  %s", cs, err)
		}
		b.choices = append(b.choices, c)
	}
}

func readNumLine(line string) []*num {
	var row []*num
	nums := strings.Fields(line)
	for _, ns := range nums {
		n, err := strconv.Atoi(ns)
		if err != nil {
			log.Fatalf("Error converting number '%s' to int:  %s", ns, err)
		}
		row = append(row, &num{n, false})
	}
	return row
}

func (b *bingo) readcard(scanner *bufio.Scanner) bool {
	c := &card{}
	appended := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return true
		}
		if !appended {
			appended = true
			b.cards = append(b.cards, c)
		}
		c.numbers = append(c.numbers, readNumLine(line))

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
	return false
}

func readdata(input string) *bingo {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	b := &bingo{}
	scanner := bufio.NewScanner(f)
	b.readchoices(scanner)

	for b.readcard(scanner) {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	return b
}
