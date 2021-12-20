package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

type bincode string

var versions, packages int

func BincodeFromHex(hex string) (s bincode) {
	for _, c := range hex {
		switch c {
		case '0':
			s += "0000"
		case '1':
			s += "0001"
		case '2':
			s += "0010"
		case '3':
			s += "0011"
		case '4':
			s += "0100"
		case '5':
			s += "0101"
		case '6':
			s += "0110"
		case '7':
			s += "0111"
		case '8':
			s += "1000"
		case '9':
			s += "1001"
		case 'A':
			s += "1010"
		case 'B':
			s += "1011"
		case 'C':
			s += "1100"
		case 'D':
			s += "1101"
		case 'E':
			s += "1110"
		case 'F':
			s += "1111"
		}
	}
	return s
}

func (b bincode) Int() (i int) {
	if i, err := strconv.ParseInt(string(b), 2, 64); err != nil {
		log.Fatalf("cannot read int value from '%s': %v ", b, err)
		return -1
	} else {
		return int(i)
	}
}

func (b bincode) version() int {
	if len(b) > 2 {
		return b[0:3].Int()
	} else {
		return 0
	}
}

func (b bincode) typeID() int {
	if len(b) >= 6 {
		return b[3:6].Int()
	} else {
		return -1
	}
}

func (b bincode) lengthTypeID() int {
	if len(b) > 6 {
		if b[6] == '0' {
			return 0
		} else {
			return 1
		}
	} else {
		return -1
	}
}

func (b bincode) literal() (value int, nextpos int) {
	wordlength := 4 + 1
	bc := b[6:]
	resbin := bincode("")
	for i := 0; i <= len(bc)-5; i = i + wordlength {
		resbin += bc[i+1 : i+wordlength]
		nextpos += wordlength
		if bc[i] == '0' {
			break
		}
	}
	//fmt.Println("Found literal: ", resbin.Int())
	return resbin.Int(), nextpos + 6
}

func (b bincode) OperatorLength() int {
	switch b[6] {
	case '0':
		return 15
	default:
		return 11
	}
}

func (b bincode) ops15() (int, int) {
	packetlen := 15
	nextpos := 7 + packetlen
	if len(b) > 7+packetlen {
		bc := b[7 : 7+packetlen]
		lengthOfPackage := bc.Int()
		v, _ := b[nextpos : nextpos+lengthOfPackage].value()
		return v, nextpos + lengthOfPackage
	}
	return 0, math.MaxInt / 2
}

func (b bincode) ops11() (int, int) {
	packetlen := 11
	nextpos := 7 + packetlen
	v := 0
	n := 0
	if len(b) > 7+packetlen {
		bc := b[7 : 7+packetlen]
		numberOfSubPackages := bc.Int()
		for i := 0; i < numberOfSubPackages; i++ {
			v, n = b[nextpos:].value()
			nextpos += n
		}
		//nextpos := 7 + 11
		return v, nextpos
	}
	return v, math.MaxInt / 2
}

func (b bincode) operator() (value int, nextpos int) {
	if len(b) > 6 {
		switch b.lengthTypeID() {
		case 0:
			v, n := b.ops15()
			value += v
			nextpos += n
		default:
			v, n := b.ops11()
			value += v
			nextpos += n
		}
		return value, nextpos
	}
	return 0, math.MaxInt / 2
}

func (b bincode) value() (int, int) {

	result := 0
	nextpos := 0
	for nextpos < len(b) {
		typeID := b[nextpos:].typeID()
		packageVisitor(b[nextpos:])
		switch typeID {
		case -1:
			return result, nextpos
		case 4:
			v, n := b[nextpos:].literal()
			result += v
			nextpos += n
		default:
			v, n := b[nextpos:].operator()
			result += v
			nextpos += n
		}
	}
	return result, nextpos
}

func packageVisitor(b bincode) {
	packages++
	versions += b.version()
	//fmt.Printf("found package: [%s - %d %d] - #packages: %d   Sum version: %d \n", string(b), b.version(), b.typeID(), packages, versions)
}

func task1(code string) (result int) {

	versions = 0
	packages = 0
	bin := BincodeFromHex(code)
	//fmt.Printf("\n\nStarting with Code: %s  Binär: %s\n", code, bin)
	bin.value()
	//fmt.Printf("Finally Versions: %d   Packages: %d \n", versions, packages)
	return versions
}

func task2(binaries string) (result int) {
	return 230
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t : %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t : %d \n", time.Since(start), result)

}
