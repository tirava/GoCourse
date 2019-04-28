// Homework-2: Arrays, Slices, Maps, Structures, JSON
// Author: Eugene Klimov
// Date: 28 april 2019

package main

import "fmt"

// 1
// machine - common properties
type machine struct {
	brand         string
	year          int
	volume        float64
	isEngineStart bool
	isWindowsOpen bool
	fuelVolume    float64
}

// car - passenger vehicle
type car struct {
	machine
	passengerSeats int
	sport          bool
}

// truck - cargo vehicle
type truck struct {
	machine
	bodyType  string
	numOfAxes int
}

// 2
var carZhiguli = car{
	machine: machine{
		brand:         "VAZ-2105",
		year:          1994,
		volume:        1.5,
		isEngineStart: false,
		isWindowsOpen: true,
		fuelVolume:    3,
	},
	passengerSeats: 4,
	sport:          false,
}

var truckKamaz = truck{
	machine: machine{
		brand:         "KAMAZ",
		year:          2005,
		volume:        10.85,
		isEngineStart: true,
		isWindowsOpen: false,
		fuelVolume:    250,
	},
	bodyType:  "dump",
	numOfAxes: 3,
}

// 3
type queue []int

func main() {

	// 2
	carZhiguli.print()
	truckKamaz.print()
	fmt.Println("----- Go! -----")

	carZhiguli.isEngineStart = true
	carZhiguli.isWindowsOpen = false
	carZhiguli.fuelVolume = 40
	carZhiguli.sport = true

	truckKamaz.isWindowsOpen = true
	truckKamaz.fuelVolume = 5
	truckKamaz.isEngineStart = false

	carZhiguli.print()
	truckKamaz.print()

	// 3
	iQueue := make(queue, 0)
	iQueue = iQueue.Push(111)
	iQueue = iQueue.Push(222)
	iQueue = iQueue.Push(333)
	fmt.Println(iQueue)

	for _, item := range iQueue {
		iQueue, item = iQueue.Shift()
		fmt.Println(iQueue, "->", item)
	}
}

// Push is for pushing item into queue
func (queue queue) Push(item int) queue {
	return append(queue, item)
}

// Shift is for return and shifting queue items
func (queue queue) Shift() (queue, int) {
	return queue[1:], queue[0]
}

// printMachine prints common machine properties
func (mach machine) printMachine() {
	fmt.Println("Machine brand:", mach.brand)
	fmt.Println("Year of release:", mach.year)
	fmt.Println("Volume of engine", mach.volume)
	fmt.Println("Engine is started?", mach.isEngineStart)
	fmt.Println("Windows are open?", mach.isWindowsOpen)
	fmt.Println("Fuel volume:", mach.fuelVolume)
}

// printCar print Car properties
func (car car) print() {
	car.printMachine()
	fmt.Println("Number of passengers:", car.passengerSeats)
	fmt.Println("Sport mode:", car.sport)
	fmt.Println("")
}

// printTruck print Truck properties
func (truck truck) print() {
	truck.printMachine()
	fmt.Println("Type of body:", truck.bodyType)
	fmt.Println("Number of axes:", truck.numOfAxes)
	fmt.Println("")
}
