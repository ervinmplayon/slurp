package main

import (
	"fmt"
	"time"
)

// * returns a function that accepts another function as an argument
func rateLimiter(interval time.Duration) func(func()) {
	var lastTime time.Time

	// * the function closure
	return func(fn func()) {
		now := time.Now()
		if now.Sub(lastTime) >= interval {
			lastTime = now
			// * call the user-provided function, execute the logic that was passed in.
			fn()
		} else {
			fmt.Println("Rate limit exceeded. Skiiping call.")
		}
	}
}

func main() {
	limit := rateLimiter(2 * time.Second)

	for i := 0; i < 5; i++ {
		/*
		 * func() {...}, is the argument passed into the rate limiter.
		 * Inside the limiter is what runs that Println.
		 * It's deferring the execution of a block of logic, based on timing
		 */
		limit(func() {
			fmt.Println("Function called at", time.Now().Format("15:04:05"))
		})
		time.Sleep(1 * time.Second)
	}
}
