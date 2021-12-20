package main

import (
	"fmt"
	"log"
)

type InfiniteArray struct {
	data             [][]interface{}
	dimX, dimY       int
	defaultValue     interface{}
	elementFormatter string
}

func NewInfiniteArray(dimX, dimY int) *InfiniteArray {
	a := &InfiniteArray{
		data: make([][]interface{}, dimY),
		dimX: dimX,
		dimY: dimY,
	}
	for i := range a.data {
		a.data[i] = make([]interface{}, dimX)
	}

	return a
}

func NewInfiniteArrayFromString(s []string, defaultValue interface{}) *InfiniteArray {
	if len(s) > 0 {
		a := &InfiniteArray{}
		a.dimY = len(s)
		a.dimX = len(s[0])
		a.defaultValue = defaultValue
		a.elementFormatter = "%c"

		for _, line := range s {
			l := make([]interface{}, len(line))
			a.data = append(a.data, l)
			for i, c := range line {
				l[i] = c
			}
		}
		return a
	}
	log.Fatal("NewInfiniteArrayFromString: cannot create from empty string")
	return nil
}

func (a *InfiniteArray) String() string {
	s := ""
	for _, line := range a.data {
		for _, c := range line {
			s += fmt.Sprintf(a.elementFormatter, c)
		}
		s += "\n"
	}
	return s
}

func (a *InfiniteArray) Get(x, y int) interface{} {
	if a.dimX > x && a.dimY > y && x >= 0 && y >= 0 {
		return a.data[y][x]
	}
	return a.defaultValue
}

func (a *InfiniteArray) GetInt(x, y int) int {
	return a.Get(x, y).(int)
}

func (a *InfiniteArray) GetBool(x, y int) bool {
	return a.Get(x, y).(bool)
}

func (a *InfiniteArray) GetString(x, y int) string {
	return fmt.Sprintf(a.elementFormatter, a.Get(x, y))
}

func (a *InfiniteArray) GetRune(x, y int) rune {
	return a.Get(x, y).(rune)
}
func (a *InfiniteArray) GetByte(x, y int) byte {
	return a.Get(x, y).(byte)
}
