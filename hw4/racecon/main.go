package main

import (
	"fmt"
	"sync"
	"time"
)

var cnt int

func incr(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		val := cnt
		val++
		time.Sleep(time.Millisecond)
		cnt = val
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incr(&wg)
	go incr(&wg)

	wg.Wait()

	fmt.Println("CNT:", cnt)
}
