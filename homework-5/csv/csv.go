// Homework-5: Standard library - Part 1
// Exercise 3 - CSV example
// Author: Eugene Klimov
// Date: 09 may 2019
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Common constants
const (
	CsvFile    = "./users.csv"
	CsvFileNew = "./users_new.csv"
	CommaIn    = ','
	CommaOut   = '\t'
)

// Person is a common person scruct
type Person struct {
	firstName    string
	lastName     string
	accountLogin string
	emailAddr    string
	mobPhone     string
}

// CSV example reads existing csv file into slice
// Prints it
// Changes some items in it
// And writes into new csv-file with other delimiter
func main() {

	// get csv data from file
	csvFile, err := os.Open(CsvFile)
	check(err)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = CommaIn
	csvData, err := reader.ReadAll()
	check(err)

	log.Println("File", CsvFile, "read successfully!")

	// fill persons from csv data
	persons := make([]Person, 0, len(csvData))
	slice2CSV2Slice(&csvData, &persons, false)

	log.Println("Persons filled...")

	// print persons
	fmt.Printf("%v\n", csvData[0]) //header
	fmt.Println(persons)           // persons data in custom string

	// change some items
	persons[0].firstName, persons[0].lastName = "*Eugen", "*Klimov"
	persons[5].emailAddr = "*------------------"
	persons[9].firstName, persons[9].lastName, persons[9].mobPhone = "*John", "*Klim", "*+12345678910"
	slice2CSV2Slice(&csvData, &persons, true)

	log.Println("Persons edited...")

	// print changed persons
	fmt.Println(persons) // persons data in custom string

	// write new file from changed csv data
	newFile, err := os.Create(CsvFileNew)
	check(err)
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.Comma = CommaOut
	err = writer.WriteAll(csvData)
	check(err)

	log.Println("File", CsvFileNew, "written successfully!")
}

//slice2CSV2Slice fills csv from persons slice and vise versa
func slice2CSV2Slice(csv *[][]string, p *[]Person, reverse bool) {
	for i, column := range *csv {
		if i == 0 {
			continue // skip header
		}
		if !reverse {
			p1 := Person{
				firstName:    column[0],
				lastName:     column[1],
				accountLogin: column[2],
				emailAddr:    column[3],
				mobPhone:     column[4],
			}
			*p = append(*p, p1)
		} else {
			column[0] = (*p)[i-1].firstName
			column[1] = (*p)[i-1].lastName
			column[2] = (*p)[i-1].accountLogin
			column[3] = (*p)[i-1].emailAddr
			column[4] = (*p)[i-1].mobPhone
		}
	}
}

// String needs for override the built-in string Person struct
func (p Person) String() string {
	f := "\n %s\t%s\t%s\t%s\t%s"
	s := fmt.Sprintf(f, p.firstName, p.lastName, p.accountLogin, p.emailAddr, p.mobPhone)
	return s
}

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
