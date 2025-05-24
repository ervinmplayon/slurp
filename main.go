package main

import (
	"fmt"
	"time"
)

// return a function that only allows execution once every `interval`
func rateLimiter(interval time.Duration) func(func()) {
	var lastTime time.Time

	return func(fn func()) {
		now := time.Now()
		if now.Sub(lastTime) >= interval {
			lastTime = now
			fn()
		} else {
			fmt.Println("Rate limit exceeded. Skiiping call.")
		}
	}
}

func main() {
	limit := rateLimiter(2 * time.Second)

	for i := 0; i < 5; i++ {
		limit(func() {
			fmt.Println("Function called at", time.Now().Format("15:04:05"))
		})
		time.Sleep(1 * time.Second)
	}
}
