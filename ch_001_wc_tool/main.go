package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var shouldCountBytes bool
	var shouldCountLines bool

	flag.BoolVar(&shouldCountBytes, "c", false, "Counts number of bytes in a file")
	flag.BoolVar(&shouldCountLines, "l", false, "Counts number of line in a file")
	flag.Parse()
	files := flag.Args()

	if shouldCountBytes {
		countBytes(files)
	}
	if shouldCountLines {
		countLines(files)
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

		fh, err := os.Open(f)
		if err != nil {
			panic(err)
		}
		sc := bufio.NewScanner(fh)
		line := 0
		for sc.Scan() {
			line++
		}

		total += line
		fmt.Println(fmt.Sprintf("%d %s", line, f))
	}
	if len(files) > 1 {
		fmt.Println(fmt.Sprintf("%d total", total))
	}
	return nil
}

func countWords(files []string) error {
	total := 0
	for _, f := range files {

		content, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		wc := 0
		var s bool
		var pr bool
		for _, b := range content {
			if (b == 10 || b == 32) && !pr{
				wc++
				pr = true
			} else {
				s = true
				pr = false
			}
		}

		if s {
			wc++
		}
		total += wc
		fmt.Println(fmt.Sprintf("%d %s", wc, f))
	}
	if len(files) > 1 {
		fmt.Println(fmt.Sprintf("%d total", total))
	}
	return nil
}
