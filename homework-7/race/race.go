// Homework-7: Goroutines & Channels
// Exercise 5 - Race emulation
// Author: Eugene Klimov
// Date: 23 may 2019
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numOfCars = 3

type car struct {
	num   int           // car number
	wait  time.Duration // car start ready time
	speed int           // car speed
}

var wg sync.WaitGroup

func main() {
	cars := make([]car, numOfCars)
	rand.Seed(time.Now().UnixNano())

	// fill start & speed
	// add wait groups
	// go all cars
	for i := 0; i < numOfCars; i++ {
		cars[i] = car{
			num:   i + 1,
			wait:  time.Duration(rand.Intn(10)),
			speed: rand.Intn(180),
		}
		wg.Add(1)
		fmt.Println("Car #", i+1, " is waiting for start", int(cars[i].wait), "seconds...")
		go cars[i].waitMe()
	}

	//wait start for all cars
	wg.Wait()

}

// waitCar method waits car for start
func (c car) waitMe() {
	defer wg.Done()

	time.Sleep(c.wait * time.Second)
	fmt.Println("Car #", c.num, "is ready for start!")
}
