package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// select only unique values
func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.RWMutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0..9
	}
	// 3, 4, 5, 0, 4, 9, 9, 8, 6, 6, 5, 5, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)

	for i := range doubles {

		go func() {
			mu.RLock()
			_, ok := alreadyStored[doubles[i]]
			mu.RUnlock()
			if !ok {
				mu.Lock()
				alreadyStored[doubles[i]] = struct{}{}
				mu.Unlock()

				uniqueIDs <- doubles[i]
			}
			if i == len(doubles)-1 {
				close(uniqueIDs)
			}
		}()

	}

	for val := range uniqueIDs {
		fmt.Println(val)
	}

	fmt.Println(uniqueIDs)
}
