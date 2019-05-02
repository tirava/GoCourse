// Homework-4: Methods and Interfaces
// Author: Eugene Klimov
// Date: 03 may 2019
package main

import (
	"./addressbook"
	"./vehicle"
	"fmt"
	"sort"
)

// vehMap needs for store vehicles map and optimized printing
type vehMap map[string]*vehicle.Vehicler

func main() {

	// 1
	fmt.Println("")

	// Zhiguli & Velik are examples of bad names of variables
	veh := make(vehMap, 2)
	var Zhiguli vehicle.Vehicler = new(vehicle.Car)
	veh["Zhiguli"] = &Zhiguli
	var Velik vehicle.Vehicler = new(vehicle.Bike)
	veh["Velik"] = &Velik

	vehicle.Buy(Zhiguli, 3456.78)
	vehicle.List(Zhiguli)
	vehicle.Buy(Velik, 1234.56)
	vehicle.List(Velik)

	fmt.Printf("Overall price: %.2f\n\n", vehicle.GetSumPrices(Zhiguli, Velik))

	fmt.Printf("----- Go! -----\n\n")

	for s, v := range veh {
		if vehicle.Go(*v) {
			fmt.Printf(s + " Goes successfully )\n\n")
		} else {
			fmt.Printf(s + " fails to Go, repairing... (\n")
			vehicle.Repair(*v, true)
			fmt.Printf("...ok!\n\n")
		}
		vehicle.List(*v)
	}

	fmt.Printf("Overall price after Go: %.2f\n\n", vehicle.GetSumPrices(Zhiguli, Velik))

	// 2
	fmt.Println("----------")
	human := []addressbook.Person{
		{"John Go", 24, []string{"+70001112233", "+56734512334"}, "Moscow"},
		{"Valery Moon", 45, []string{"+79991234567"}, "Tula"},
		{"Klim Sangin", 18, []string{"+71112223344", "+15554443322", "+23458887766"}, "Minsk"},
		{"Kir Korov", 30, []string{"+800987654321"}, "Aldebaran"},
	}

	sort.Sort(addressbook.ByAge(human))
	fmt.Println("\nSorted by Age:")
	printAddressBook(&human)

	sort.Sort(addressbook.ByAddr(human))
	fmt.Println("\nSorted by Address:")
	printAddressBook(&human)
}

// printAddressBook prints my address book more friendly
func printAddressBook(ab *[]addressbook.Person) {
	for _, item := range *ab {
		fmt.Printf("\nName: %v\n", item.Name)
		fmt.Printf("\tAge: %v\n", item.Age)
		for i, number := range item.Phones {
			fmt.Printf("\tPh.%v: %v \n", i+1, number)
		}
		fmt.Printf("\tAddress: %v\n", item.Address)
	}
}
