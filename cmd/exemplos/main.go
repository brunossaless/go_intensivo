package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for num := range data {
		fmt.Printf("Worker %d received %d\n", workerId, num)
		time.Sleep(time.Second)
	}
}

func main() { // gorountine 1

	ch := make(chan int)
	qtdWorkers := 3

	//create workers and go routines
	for i := range qtdWorkers {
		go worker(i+1, ch)
	}

	for i := range 10 {
		ch <- i
	}
}
