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
	fmt.Println(carZhiguli)
	fmt.Println(truckKamaz)

	carZhiguli.isEngineStart = true
	carZhiguli.isWindowsOpen = false
	carZhiguli.fuelVolume = 40
	carZhiguli.sport = true
	fmt.Println(carZhiguli)

	truckKamaz.isWindowsOpen = true
	truckKamaz.fuelVolume = 5
	truckKamaz.isEngineStart = false
	fmt.Println(truckKamaz)

}
