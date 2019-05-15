// Homework-6: Standard library - Part 2
// Exercise 4 - Square Equation
// Author: Eugene Klimov
// Date: 16 may 2019
package main

import (
	"fmt"
	"log"
	"net/http"
)

// hostAndPort for listen
const hostAndPort = "localhost:8080"

// Open http://hostAndPort enter parameters and click Send
// Or directly enter http://hostAndPort/hello?name=Eugene&age=25
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	fmt.Println("Open URL in your browser: http://" + hostAndPort)
	err := http.ListenAndServe(hostAndPort, nil)
	check(err, "fatal", "Can't start server!")

}

// hello handler
func hello(res http.ResponseWriter, req *http.Request) {

	// get form keys
	name := req.FormValue("name")
	age := req.FormValue("age")

	// write to response
	fmt.Fprintf(res, "Your name is: %s\nYour age is: %s", name, age)
}

// check is errors helper
func check(err error, sType, sMessage string) {
	if err != nil {
		switch sType {
		case "fatal":
			log.Fatalln(sMessage, err)
		default:
			log.Println(sMessage, err)
		}
	}
}
