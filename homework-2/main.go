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
	var num int
	fmt.Println("Enter integer number:")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal(err)
	}
	if isNumEven(num) {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 2

}

// isNumEven checks nums for even/odd
func isNumEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
