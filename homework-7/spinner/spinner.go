// Homework-7: Goroutines & Channels
// Exercise 1 - Spinner for 10 seconds
// Author: Eugene Klimov
// Date: 21 may 2019
package main

import (
	"fmt"
	"time"
)

// Spinner demo in goroutine
func main() {
	go spinner(50 * time.Millisecond)
	//	const n = 45
	//	fibN := fibonacci(n)
	//	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	time.Sleep(10 * time.Second) // need spin for 10 seconds only
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

/*func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}*/
