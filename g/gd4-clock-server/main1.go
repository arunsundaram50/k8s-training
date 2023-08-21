package main

import (
	"fmt"
	"time"
)

func serveTime(out chan<- string) {
	for {
		out <- time.Now().Format(time.UnixDate)
		time.Sleep(time.Second)
	}
}

func printTime(inChannel <-chan string) {
	for timeString := range inChannel {
		fmt.Printf("%s\n", timeString)
	}
}

func main() {
	timeChannel := make(chan string)
	go serveTime(timeChannel)
	printTime(timeChannel)
}
