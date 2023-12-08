package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		value := <-ch
		fmt.Println("Goroutine1:", value)
		ch <- 42
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 23
		value := <-ch
		fmt.Println("Goroutine2:", value)
	}()

	wg.Wait()
}
