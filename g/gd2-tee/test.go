package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tw := NewTeeWriter(file, os.Stdout)
	_, err = tw.Write([]byte("Hello, world1\n"))
	if err != nil {
		fmt.Println(err)
	}
}
