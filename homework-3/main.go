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

func main() {

	// 2
	carZhiguli.printCar()
	truckKamaz.printTruck()
	fmt.Println("----- Go! -----")

	carZhiguli.isEngineStart = true
	carZhiguli.isWindowsOpen = false
	carZhiguli.fuelVolume = 40
	carZhiguli.sport = true

	truckKamaz.isWindowsOpen = true
	truckKamaz.fuelVolume = 5
	truckKamaz.isEngineStart = false

	carZhiguli.printCar()
	truckKamaz.printTruck()
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
func (car car) printCar() {
	car.printMachine()
	fmt.Println("Number of passengers:", car.passengerSeats)
	fmt.Println("Sport mode:", car.sport)
	fmt.Println("")
}

// printTruck print Truck properties
func (truck truck) printTruck() {
	truck.printMachine()
	fmt.Println("Type of body:", truck.bodyType)
	fmt.Println("Number of axes:", truck.numOfAxes)
	fmt.Println("")
}
