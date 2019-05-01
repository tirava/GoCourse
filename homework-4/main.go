// Homework-4: Methods and Interfaces
// Author: Eugene Klimov
// Date: 02 may 2019

package main

import (
	"./vehicle"
	"fmt"
)

//type addressBook map[string][]string

//const jsonName = "ab.json"

func main() {

	// Zhiguli & Velik are examples of bad names of variables
	var Zhiguli vehicle.Vehicler = new(vehicle.Car)
	var Velik vehicle.Vehicler = new(vehicle.Bike)

	vehicle.Buy(Zhiguli, 3456.78)
	vehicle.List(Zhiguli)
	vehicle.Buy(Velik, 1234.56)
	vehicle.List(Velik)

	fmt.Printf("----- Go! -----\n\n")

	if vehicle.Go(Zhiguli) {
		fmt.Println("Zhiguli Goes succsessfully )")
		vehicle.List(Zhiguli)
	} else {
		fmt.Printf("Zhiguli fails to Go (\n\n")
	}

	if vehicle.Go(Velik) {
		fmt.Println("Velik Goes succsessfully )")
		vehicle.List(Velik)
	} else {
		fmt.Printf("Velik fails to Go (\n\n")
	}

	// 2
	/*ab := make(addressBook, 2)
	ab["John"] = []string{"+70001112233"}
	ab["Klim"] = []string{"+71112223344"}
	ab["Klim"] = append(ab["Klim"], "+15554443322")
	fmt.Println("")
	printAddressBook(&ab)

	f, err := saveJsonAB(&ab)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nAddressBook saved as", f)*/
}

// saveJsonAB saves Address Book as JSON file
/*func saveJsonAB(ab *addressBook) (path string, err error) {
	b, err := json.Marshal(*ab)
	if err != nil {
		return "", err
	}

	filePath := filepath.Dir(os.Args[0]) // for correct path run "go build" instead "go run"
	fileName := filepath.Join(filePath, jsonName)

	err = ioutil.WriteFile(fileName, b, 0777)

	return fileName, err
}*/

// printAddressBook prints any Address Book
/*func printAddressBook(ab *addressBook) {
	for name, numbers := range *ab {
		fmt.Println("Abonent:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v: %v \n", i+1, number)
		}
	}
}*/
