package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{ //"example0.txt", 2021, 230},
	{"example1.txt", 9, 230},
	{"example2.txt", 14, 230},
	{"example3.txt", 16, 230},
	{"example4.txt", 12, 230},
	{"example5.txt", 23, 230},
	{"example6.txt", 31, 230}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		input := readdata(test.fname)
		r := task1(input)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		input := readdata(test.fname)
		r := task2(input)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
