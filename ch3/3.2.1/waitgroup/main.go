package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine sleeping...")
	// 	time.Sleep(1 * time.Second)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd goroutine sleeping...")
	// 	time.Sleep(2 * time.Second)
	// }()

	// wg.Wait()
	// fmt.Println("All gorutines complete.")

	/////////////////////////////////////////////////
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
