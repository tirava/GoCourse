// Syntax Go:	Homework-2
// Author:		Eugene Klimov
// Date:		23 april 2019

package main

import (
	"fmt"
	"log"
	"strconv"
)

const maxFibonacci = 93
const maxPrime = 500
const maxPrimeIndex = 3571 + 1 // max integer for prime = 500 (see wiki)

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

	// 3 & 4
	msg = "\nEnter amount of the first Fibonacci numbers (min = 2, max = " + strconv.Itoa(maxFibonacci) + "):"
	for {
		num = getIntNum(&msg)
		if num > maxFibonacci || num < 2 {
			continue
		} else {
			break
		}
	}
	fmt.Println("First "+strconv.Itoa(num)+" Fibonacci numbers:\n", getNFibonacci(num, 0))

	// 5
	msg = "\nEnter amount of elements for the prime numbers slice (min = 1, max = " + strconv.Itoa(maxPrime) + "):"
	for {
		num = getIntNum(&msg)
		if num > maxPrime || num < 1 {
			continue
		} else {
			break
		}
	}
	fmt.Println("First "+strconv.Itoa(num)+" prime numbers:\n", fillPrimeNumber(num))
}

// fillPrimeNumber fills slice contains N elements for the prime numbers
func fillPrimeNumber(n int) (slice []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Don't panic, everything is under control )", r)
		}
	}()

	boolArr := make([]bool, maxPrimeIndex) //max prime index for first 500 numbers
	boolArr[0] = true                      // non prime
	boolArr[1] = true                      // non prime

	// fill bool array non prime numbers, prime = false for default
	for i := 2; i < maxPrimeIndex; i++ {
		if !boolArr[i] {
			if i*i < maxPrimeIndex {
				for j := i * i; j < maxPrimeIndex; j += i {
					boolArr[j] = true
				}
			}
		}
	}
	// convert bool indexes to prime numbers
	primeArr := make([]int, 0, n)
	for i := 0; i < maxPrimeIndex; i++ {
		if !boolArr[i] {
			primeArr = append(primeArr, i)
			n--
			if n == 0 {
				break // no need more than n
			}
		}
	}
	return primeArr
}

// getNFibonacci returns N first Fibonacci numbers in slice, first = 0 or 1
func getNFibonacci(n int, first int) (slice []uint64) {
	arr := make([]uint64, n)
	arr[0] = uint64(first)
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
		// uncomment to print vertical
		// println(arr[i])
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
