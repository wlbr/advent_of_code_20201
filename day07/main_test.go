package main

import (
	"testing"
)

type testdata struct {
	fname     string
	expected1 int
	expected2 int
}

var testset []*testdata = []*testdata{{"example.txt", 37, 168}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		input := readdata(test.fname)
		_, c := task1(input)
		if c != test.expected1 {
			t.Fatalf("TestTaskOne '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expected1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		input := readdata(test.fname)
		_, c := task2(input)
		if c != test.expected2 {
			t.Fatalf("TestTaskTwo '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expected2)
		}
	}
}
