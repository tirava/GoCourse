// Homework-7: Goroutines & Channels
// Exercise 2 - Pipeline
// Author: Eugene Klimov
// Date: 21 may 2019
package main

import (
	"fmt"
	"time"
)

// Pipeline for fixed values
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for x := range squares {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
