package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Person struct {
	firstName    string
	lastName     string
	accountLogin string
	emailAddr    string
	mobPhone     string
}

// csvExamples reads existing csv file
// Prints it
// Changes some items in it
// And writes into new csv-file with other delimiter
func csvExamples() {

	csvFile, err := os.Open("./users.csv")
	check(err)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ','

	csvData, err := reader.ReadAll()
	check(err)

	// write to struct
	persons := make([]Person, 0, len(csvData))
	for i, column := range csvData {
		if i == 0 {
			continue // skip header
		}
		p := Person{
			firstName:    column[0],
			lastName:     column[1],
			accountLogin: column[2],
			emailAddr:    column[3],
			mobPhone:     column[4],
		}
		persons = append(persons, p)
	}

	fmt.Printf("%v\n", csvData[0]) //header
	fmt.Println(persons)           // csv data in custom string
}

// String is for override the built-in string Person struct
func (p Person) String() string {
	f := "\n %s\t%s\t%s\t%s\t%s"
	s := fmt.Sprintf(f, p.firstName, p.lastName, p.accountLogin, p.emailAddr, p.mobPhone)
	return s
}
