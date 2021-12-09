package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 0, 5353}, {"example2.txt", 26, 61229}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		signals, patterns := readdata(test.fname)
		r := task1(signals, patterns)
		if r != test.expectedtask1 {
			t.Fatalf("TestTaskOne '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset[:1] {
		signals, patterns := readdata(test.fname)
		r := task2(signals, patterns)
		if r != test.expectedtask2 {
			t.Fatalf("TestTaskTwo '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}

func TestLogic(t *testing.T) {

	a := "abcd"
	not_expected := "efg"
	b := "bcde"
	//c := "efg"

	and_expected := "bcd"
	or_expected := "abcde"
	xor_expected := "ae"

	if res := not(a); res != not_expected {
		t.Fatalf("TestLogic 'not' failed. Got '%s' -  Wanted: '%s'", res, not_expected)
	}
	if res := or(a, b); res != or_expected {
		t.Fatalf("TestLogic 'OR' failed. Got '%s' -  Wanted: '%s'", res, or_expected)
	}
	if res := xor(a, b); res != xor_expected {
		t.Fatalf("TestLogic 'XOR' failed. Got '%s' -  Wanted: '%s'", res, xor_expected)
	}
	if res := and(a, b); res != and_expected {
		t.Fatalf("TestLogic 'AND' failed. Got '%s' -  Wanted: '%s'", res, and_expected)
	}
}
