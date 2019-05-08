package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	bs, err := ioutil.ReadFile("filereadshort.go")
	// if err != nil {
	// Лучше логгировать в консоль и выходить
	// return
	// }
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
