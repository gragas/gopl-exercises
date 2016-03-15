// Exercise 1.3: Experiment to measure the difference in running
// time between our potentially inefficient versions and the one
// that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark
// tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"strings"
	"time"
)

const NumArgs = 80000

func main() {
	// Generate fake arguments
	var fakeArgs [NumArgs]string
	for i := 0; i < NumArgs; i++ {
		fakeArgs[i] = "a"
	}

	// Time the naive implementation
	fmt.Printf("Timing %d arguments (naive implementation)...", NumArgs)
	start := time.Now()
	var s, sep string
	for _, arg := range(fakeArgs) {
		s += arg + sep
		sep = " "
	}
	fmt.Printf("%.2f seconds\n", time.Since(start).Seconds())

	// Time the string.Joins implementation
	fmt.Printf("Timing %d arguments (strings.Join)...", NumArgs)
	start = time.Now()
	strings.Join(fakeArgs[:], " ")
	fmt.Printf("%.2f seconds\n", time.Since(start).Seconds())	
}
