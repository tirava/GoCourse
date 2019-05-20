// Homework-7: Go routines & channels
// Exercise 1 - Fibonacci numbers off
// Author: Eugene Klimov
// Date: 21 may 2019
package main

import (
	"fmt"
	"time"
)

// demo fibonacci off in go routines
func main() {

	go spinner(50 * time.Millisecond)

	//	const n = 45
	//	fibN := fibonacci(n)

	time.Sleep(10 * time.Second) // need sleep for 10 seconds only

	//	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
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
