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
	// Cursor display concerns in one block
	defer func() func() {
		fmt.Print("\033[?25l") // Hide
		go spinner()           // go makes spinner() return so that we can carry on
		return func() {
			fmt.Print("\033[?25h") // Show
		}
	}()()

	// Validation
	if len(os.Args) != 2 {
		panic("Need one int parameter to calculate its slow fibonacci value")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Computation
	var n = slowFib(i)
	fmt.Printf("\nFibonacci value of %d is %d\n", i, n)
}
