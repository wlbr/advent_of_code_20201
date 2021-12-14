package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 1588, 2188189693529}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		p, r := readdata(test.fname)
		res := task1(p, r)
		if res != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, res, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		p, r := readdata(test.fname)
		res := task2(p, r)
		if res != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, res, test.expectedtask2)
		}
	}
}
