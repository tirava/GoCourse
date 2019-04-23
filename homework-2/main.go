// Syntax Go:	Homework-2
// Author:		Eugene Klimov
// Date:		23 april 2019

package main

import (
	"fmt"
	"log"
)

func main() {

	// 1
	msg := "\nEnter integer number fo check even/odd:"
	num := getIntNum(&msg)
	if isNumDivisibleBy(num, 2) {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 2
	msg = "\nEnter integer number for check divisible by 3:"
	num = getIntNum(&msg)
	if isNumDivisibleBy(num, 3) {
		fmt.Println("The number is divisible by 3")
	} else {
		fmt.Println("The number is not divisible by 3")
	}
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
func isNumDivisibleBy(num int, div int) bool {
	if num%div == 0 {
		return true
	} else {
		return false
	}
}
