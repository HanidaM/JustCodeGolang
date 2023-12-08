package main

import (
	"fmt"
	"sync"
)

func megre(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	copy := func(ch <-chan int) {
		defer wg.Done()
		for value := range ch {
			merged <- value
		}
	}

	wg.Add(len(channels))

	for _, ch := range channels {
		go copy(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	merged := megre(ch1, ch2)

	for value := range merged {
		fmt.Println(value)
	}
}
