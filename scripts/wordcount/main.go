package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type kv struct {
	Key   string
	Value int
}

func main() {
	rawPath := os.Args[1]
	count := make(map[string]int)

	file, err := os.Open(rawPath)
	if err != nil {
		log.Fatalf("failed to open: %v file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count[scanner.Text()]++
	}

	var list []kv
	for k, v := range count {
		list = append(list, kv{k, v})
	}
	slices.SortFunc(list, func(a, b kv) int {
		return b.Value - a.Value
	})

	for i := range list {
		if i < 10 {
			fmt.Println(list[i].Key, list[i].Value)
		}
	}

}
