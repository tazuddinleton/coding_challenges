package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var shouldCountBytes bool

	flag.BoolVar(&shouldCountBytes, "c", false, "Counts number of bytes in a file")
	flag.Parse()
	files := flag.Args()

	if shouldCountBytes {
		countBytes(files)
	}
}

func countBytes(files []string) error {
	var total int
	for _, f := range files {
		bytes, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		n := len(bytes)
		total += n
		fmt.Println(fmt.Sprintf("%d %s", n, f))
	}
	if len(files) > 1 {
		fmt.Println(fmt.Sprintf("%d total", total))
	}
	return nil
}

func countLines(files []string) error {
	var total int
	for _, f := range files {
		bytes, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		n := len(bytes)
		total += n
		fmt.Println(fmt.Sprintf("%d %s", n, f))
	}
	if len(files) > 1 {
		fmt.Println(fmt.Sprintf("%d total", total))
	}
	return nil
}
