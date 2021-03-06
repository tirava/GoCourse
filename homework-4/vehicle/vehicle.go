// vehicle implements vehicle interface
package vehicle

import (
	"fmt"
)

type Vehicler interface {
	Buy(price float64)
	List()
	Go() bool // is successful
	GetPrice() float64
	Repair(bool)
}

type Car struct {
	kind           string
	brand          string
	year           int
	volume         float64 // engine in m^3
	maxspeed       int     // km/h
	isEngineStart  bool
	isWindowsOpen  bool
	fuelVolume     float64 // in litres
	passengerSeats int
	sport          bool    // sport mode\
	price          float64 // dollars
	serviceable    bool
}

type Bike struct {
	kind        string
	brand       string
	gender      string
	target      string
	wheels      float64
	electric    bool
	price       float64
	serviceable bool
}

// Buy interface
func Buy(v Vehicler, price float64) {
	v.Buy(price)
}

// List interface
func List(v Vehicler) {
	v.List()
}

// Go feature
func Go(v Vehicler) bool {
	return v.Go()
}

// Repair vehicle
func Repair(v Vehicler, rep bool) {
	v.Repair(rep)
}

// GetSumPrices returns sum of all prices
func GetSumPrices(price ...Vehicler) float64 {
	var sum float64
	for _, p := range price {
		sum += p.GetPrice()
	}
	return sum
}

// Buy car
func (c *Car) Buy(price float64) {
	c.kind = "Car"
	c.brand = "VAZ-2105"
	c.year = 1994
	c.volume = 1.5
	c.maxspeed = 120
	c.isEngineStart = false
	c.isWindowsOpen = true
	c.fuelVolume = 3
	c.passengerSeats = 4
	c.sport = false
	c.price = price
	c.serviceable = false
}

// Buy bike
func (b *Bike) Buy(price float64) {
	b.kind = "Bike"
	b.brand = "Merida"
	b.gender = "man"
	b.target = "Mountain"
	b.wheels = 29 // inches
	b.electric = false
	b.price = price
	b.serviceable = true
}

// List car properties - simple print but need store in the array in feature
func (c Car) List() {
	fmt.Println("Kind of vehicle:", c.kind)
	fmt.Println("Machine brand:", c.brand)
	fmt.Println("Year of release:", c.year)
	fmt.Println("Volume of engine", c.volume)
	fmt.Println("Maximum speed", c.maxspeed)
	fmt.Println("Engine is started?", c.isEngineStart)
	fmt.Println("Windows are open?", c.isWindowsOpen)
	fmt.Println("Fuel volume:", c.fuelVolume)
	fmt.Println("Number of passengers:", c.passengerSeats)
	fmt.Println("Sport mode:", c.sport)
	fmt.Println("Price:", c.price)
	fmt.Println("Serviceable:", c.serviceable)
	fmt.Println("")
}

// List bike properties - simple print but need store in the array in feature
func (b Bike) List() {
	fmt.Println("Kind of vehicle:", b.kind)
	fmt.Println("Machine brand:", b.brand)
	fmt.Println("Intended gender:", b.gender)
	fmt.Println("Target use:", b.target)
	fmt.Println("Wheels size:", b.wheels)
	fmt.Println("Use electric engine?", b.electric)
	fmt.Println("Price:", b.price)
	fmt.Println("Serviceable:", b.serviceable)
	fmt.Println("")
}

// Go starts car to go
func (c *Car) Go() bool {
	c.sport = true
	c.isWindowsOpen = false
	c.price = c.price / 2
	return c.serviceable // false - engine does not work
}

// Go starts bike to go
func (b *Bike) Go() bool {
	b.target = "Road"
	b.electric = true
	b.price = b.price / 2
	return b.serviceable // true - all is ok
}

// GetPrice returns current car price
func (c *Car) GetPrice() float64 {
	return c.price
}

// GetPrice returns current bike price
func (b *Bike) GetPrice() float64 {
	return b.price
}

// Repair need for service car
func (c *Car) Repair(rep bool) {
	c.serviceable = rep
}

// Repair need for service bike
func (b *Bike) Repair(rep bool) {
	b.serviceable = rep
}
