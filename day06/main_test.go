package main

import (
	"math/big"
	"testing"
)

type testdata struct {
	fname          string
	days           int
	expectedfishes int64
}

var testset []*testdata = []*testdata{{"example.txt", 18, 26},
	{"example.txt", 80, 5934},
	{"example.txt", 256, 26984457539}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		input := readdata(test.fname)
		r := task1and2(input, test.days)
		if r.Cmp(big.NewInt(test.expectedfishes)) != 0 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedfishes)
		}
	}
}
