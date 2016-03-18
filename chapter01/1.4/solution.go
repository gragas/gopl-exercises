// Exercise 1.4: Modify dup2 to print the names of all the files in
// which each duplicated line appears

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	relations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) < 1 {
		countLines(os.Stdin, counts, "os.Stdin", relations)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, relations)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", line, count)
			for _, f := range relations[line] {
				fmt.Printf("-> %v\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fname string, relations map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		relations[input.Text()] = append(relations[input.Text()], fname)
	}
}
