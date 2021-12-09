package main

import "strings"

func or(a string, ops ...string) string {
	newa := a
	for _, a := range ops {
		for _, c := range a {
			if !strings.Contains(newa, string(c)) {
				newa += string(c)
			}
		}
	}
	return sortString(newa)
}

func xor(a string, ops ...string) string {
	newa := a
	for _, a := range ops {
		for _, c := range a {
			if !strings.Contains(newa, string(c)) {
				newa += string(c)
			} else {
				newa = strings.Replace(newa, string(c), "", -1)
			}
		}
	}
	return sortString(newa)
}

func and(first string, second string) string {
	newa := ""
	for _, c := range second {
		if strings.Contains(first, string(c)) {
			newa += string(c)
		}
	}
	return sortString(newa)
}

func not(w string) (i string) {
	keys := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	for _, k := range keys {
		if !strings.Contains(w, string(k)) {
			i += string(k)
		}
	}
	return i
}
