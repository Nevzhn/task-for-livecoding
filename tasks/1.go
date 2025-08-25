package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// select only unique values
func main() {
	var alreadyStored map[int]struct{}
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0..9
	}
	// 3, 4, 5, 0, 4, 9, 9, 8, 6, 6, 5, 5, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				mu.Lock()
				alreadyStored[doubles[i]] = struct{}{}
				mu.Unlock()

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	for val := range uniqueIDs {
		fmt.Println(val)
	}

	fmt.Println(uniqueIDs)
}
