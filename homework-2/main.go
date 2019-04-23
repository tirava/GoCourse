// Syntax Go:	Homework-2
// Author:		Eugene Klimov
// Date:		23 april 2019

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// 1
	msg := "\nEnter integer number fo check even/odd:"
	num := getIntNum(&msg)
	result := isNumDivisibleBy(num, 2)
	if result == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 2
	msg = "\nEnter integer number for check divisible by 3:"
	num = getIntNum(&msg)
	result = isNumDivisibleBy(num, 3)
	if result == 0 {
		fmt.Println("The number is divisible by 3 without remainder")
	} else {
		fmt.Println("The number is not divisible by 3, remainder:", result)
	}

	//3
	msg = "\nEnter amount of the first Fibonacci numbers:"
	num = getIntNum(&msg)
	fmt.Println("First "+strconv.Itoa(num)+" Fibonacci numbers:\n", getNFibonacci(num, 1))
}

// getNFibonacci returns N first Fibonacci numbers in slice, first = 0 or 1
func getNFibonacci(n int, first int) (slice []int) {
	arr := make([]int, n)
	arr[0] = first
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

// getIntNum return integer number from standard input
func getIntNum(msg *string) int {
	var num int
	fmt.Println(*msg)
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// isNumDivisibleBy checks if the number without remainder divisible by div
func isNumDivisibleBy(num int, div int) int {
	result := num % div
	if result == 0 {
		return 0
	} else {
		return result
	}
}
