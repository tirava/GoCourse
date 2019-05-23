// Homework-7: Goroutines & Channels
// Exercise 5 - Race emulation
// Author: Eugene Klimov
// Date: 23 may 2019
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	numOfCars         = 5
	distance  float64 = 1   // km
	maxSpeed  float64 = 555 // km/h - if < 60 it will be fixed to 60
	maxWait           = 10  // seconds
)

type car struct {
	num   int           // car number
	wait  time.Duration // car start ready time
	speed float64       // car speed in km/h
	time  float64       // distance time
}

var wg, wgChan sync.WaitGroup
var carChan = make(chan car, numOfCars) // for fix cars parameters on finish

func main() {
	cars := make([]car, numOfCars)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nDistance:", distance, "km")
	fmt.Println("-------------------------------------------")

	// fill start & speed
	// add wait groups
	// go all cars
	for i := 0; i < numOfCars; i++ {
		cars[i] = car{
			num:   i + 1,
			wait:  time.Duration(rand.Intn(maxWait)),
			speed: rand.Float64() * maxSpeed,
		}
		if cars[i].speed < 60 { // fix for exclude div by zero and very small speeds
			cars[i].speed = 60
		}
		wg.Add(1)
		fmt.Println("Car #", i+1, "is waiting for start", int(cars[i].wait), "seconds...")
		go cars[i].waitMe()
		go spinner(50 * time.Millisecond) // without waitgroup
	}
	fmt.Println("-------------------------------------------")

	//wait start for all cars
	wg.Wait()
	fmt.Println("-------------------------------------------")

	// start all cars
	for _, car := range cars {
		fmt.Println("Car #", car.num, "is started...")
		wg.Add(1)
		go car.startMe()
	}
	fmt.Println("-------------------------------------------")

	// get times for finished cars
	// waiting need for correct calc winner (i.e clean channel)
	// but waiting may be replaced simple sleep for n seconds
	wgChan.Add(1)
	go func() {
		defer wgChan.Done()
		for _, car := range cars {
			car = <-carChan
			cars[car.num-1].time = car.time
		}
	}()

	// wait cars finished
	go spinner(50 * time.Millisecond) // without waitgroup
	wg.Wait()
	wgChan.Wait()
	fmt.Println("-------------------------------------------")

	// check winner
	winCar := 0
	minTime := math.MaxFloat64

	for _, car := range cars {
		if car.time < minTime {
			minTime = car.time
			winCar = car.num
		}
	}

	fmt.Printf("Car # %d wins the race!\n\n", winCar)
}

// waitMe method waits car for start
func (c car) waitMe() {
	defer wg.Done()

	time.Sleep(c.wait * time.Second)
	fmt.Println("Car #", c.num, "is ready for start!")
}

// startMe method starts the car for distance km and fixes finish
func (c car) startMe() {
	defer wg.Done()

	// get time in seconds for distance / speed and sleep )
	t := distance / c.speed * 3600
	time.Sleep(time.Duration(t) * time.Second)

	c.time = t
	carChan <- c // fix car parameters after finish

	fmt.Printf("Car # %d is finished for %.2f seconds! Speed: %.2f km/h\n", c.num, t, c.speed)
}

// spinner is for comfort waiting cars )
func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
