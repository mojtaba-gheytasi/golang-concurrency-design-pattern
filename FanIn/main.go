package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func send(even, odd chan<- int)  {

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}