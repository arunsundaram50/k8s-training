package main

import (
	"fmt"
	"os"
	"strconv"
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
	if len(os.Args) != 2 {
		panic("Need one int parameter to calculate its slow fibonacci value")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	go spinner()

	// Computation
	var n = slowFib(i)
	fmt.Printf("\nFibonacci value of %d is %d\n", i, n)
}
