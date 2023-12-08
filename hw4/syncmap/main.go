package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	data := make(map[string]string)

	go func() {
		mu.Lock()
		data["key"] = "val"
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		val := data["key"]
		fmt.Println("Read from map:", val)
		mu.Unlock()
	}()

	time.Sleep(time.Second)
}
