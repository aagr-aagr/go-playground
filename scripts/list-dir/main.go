package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	rawPath := os.Args[1]
	var totalSize int64
	err := filepath.WalkDir(rawPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		data, err := d.Info()
		if err != nil {
			return err
		}

		fmt.Println(d.Name())
		if d.IsDir() {
			return nil
		}
		totalSize += data.Size()
		return nil

	})

	if err != nil {
		log.Fatalf("error walking the path: %v", err)
	}

	fmt.Printf("Total size of all files: %f MB", float64(totalSize)/1000000)
}
