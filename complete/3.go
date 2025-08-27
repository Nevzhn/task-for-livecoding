package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

// вместо mutex используем атомарные числа
var counter atomic.Int64

func SimulateRequest(ctx context.Context) (int64, error) {
	t := time.Now()
	defer func() {
		log.Printf("Время выполнения запроса: %v", time.Since(t))
	}()

	ch := make(chan atomic.Int64, 1)

	go func() {
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Millisecond)
		counter.Add(1)
		ch <- counter
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("timeout request: %w", ctx.Err())
	case count := <-ch:
		return count.Load(), nil
	}
}

func main() {

	for range 100 {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
			defer cancel()

			val, err := SimulateRequest(ctx)
			if err != nil {
				log.Printf("Failed request: %v", err)
			}

			log.Printf("Значение счетчика: %d", val)
		}()
	}
}

//Требуется доработать метод SimulateRequest так, чтобы:
//- Код работал в конкурентной среде
//- При долгом ожидании запрос отваливался по таймауту
//- На консоль печаталось время выполнения запроса
