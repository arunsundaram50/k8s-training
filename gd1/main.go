package main

import (
	"fmt"
	"os"
)

func main() {
	println("args:")
	for _, arg := range os.Args {
		fmt.Printf("\t%v\n", arg)
	}
}
