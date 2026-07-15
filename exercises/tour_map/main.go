package main

import (
	"strings"
)

func wordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, y := range words {
		_, ok := result[y]
		if ok {
			result[y]++
		}
	}
	return result
}

func main() {

}
