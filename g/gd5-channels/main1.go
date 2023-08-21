package main

import (
	"fmt"
	"time"
)

func generator(out chan<- int) {
	for i := 0; ; i++ {
		out <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func squarer(outputSquareChannel chan<- int, inputNumberChannel <-chan int) {
	for i := range inputNumberChannel {
		outputSquareChannel <- i * i
	}
}

func printer(numberChannel <-chan int) {
	for i := range numberChannel {
		fmt.Printf("%d\n", i)
	}
}

func main() {
	naturalNumbersChannel := make(chan int)
	squareChannel := make(chan int)
	go generator(naturalNumbersChannel)
	go squarer(squareChannel, naturalNumbersChannel)
	printer(squareChannel)
}
