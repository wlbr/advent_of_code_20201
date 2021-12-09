package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func sortedStrings(words []string) []string {
	sorted := make([]string, len(words))
	for i, w := range words {
		sorted[i] = sortString(w)
	}
	return sorted
}

func readdata(input string) (signals, patterns [][]string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for _, word := range lines {
		parts := strings.Split(word, "|")
		if len(parts) != 2 {
			log.Fatalf("malformed input line: %s", word)
		} else {
			s := strings.Split(parts[0], " ")
			if len(s) != 11 {
				log.Fatalf("malformed signals: %s", signals)
			} else {
				signals = append(signals, sortedStrings(s[:len(s)-1]))
				p := strings.Split(parts[1], " ")
				if len(p) != 5 {
					log.Fatalf("malformed patterns: %s", patterns)
				} else {
					patterns = append(patterns, sortedStrings(p[1:]))
				}
			}
		}
	}
	return signals, patterns
}
