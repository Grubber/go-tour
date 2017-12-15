package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	array := strings.Fields(s)
	for _, i := range array {
		m[i] = m[i] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
