// Homework-5: Standard library - Part 1
// Author: Eugene Klimov
// Date: 11 may 2019
package main

import "log"

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	csvExamples()

}
