// Homework-6: Standard library - Part 2
// Exercise 4 - Square Equation
// Author: Eugene Klimov
// Date: 16 may 2019
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// hostAndPort for listen
const hostAndPort = "localhost:8080"

type helloData struct {
	PageTitle string
	Name, Age string
}

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

	if name == "" {
		name = "Евгений Климов"
	}

	if age == "" {
		age = "29"
	}

	// show template with data
	tmpl, err := template.ParseFiles("./template/hello.gohtml")
	check(err, "fatal", "Can't open template!")

	data := helloData{
		PageTitle: "Hello response",
		Name:      name,
		Age:       age,
	}

	// need to view correctly
	res.Header().Set("Content-Type", "text/html")

	err = tmpl.Execute(res, data)
	check(err, "fatal", "Can't execute template!")
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
