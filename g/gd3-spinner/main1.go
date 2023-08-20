package main

import (
	"fmt"
	"os"
	"time"
)

func spinner() {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func slowFib(n int) int {
	if n < 2 {
		return n
	}
	return slowFib(n-1) + slowFib(n-2)
}

func main() {
	// Validation
	if len(os.Args) != 1 {
		panic("I always compute slowFib(45). Sorry!")
	}

	go spinner()

	// Computation
	const N = 45
	var n = slowFib(N)
	fmt.Printf("\nFibonacci value of %d is %d\n", N, n)
}
