package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	ss := strings.Fields(s)
	for _, word := range ss {
		_, exists := m[word]
		
		if exists {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}