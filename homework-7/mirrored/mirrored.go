// Homework-7: Goroutines & Channels
// Exercise 4 - Sites mirrors
// Author: Eugene Klimov
// Date: 22 may 2019
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const fileName = "mirrored.txt" // no new line after last record!

type siteMirror struct {
	name     string
	response time.Duration
}

// Reads sites from txt file and checks more fastest
func main() {
	sitesNames := make([]siteMirror, 0)

	file, err := os.Open(fileName)
	check(err, "fatal", "Can't open file with sites!")
	defer file.Close()

	// read sites by line
	f := bufio.NewReader(file)
	for {
		line, err := f.ReadString('\n')
		if err == io.EOF {
			break
		}
		sitesNames = append(sitesNames, siteMirror{strings.TrimRight(line, "\n"), 0})
	}

	// print result
	fmt.Println("Fastest site/mirror from:")
	for _, site := range sitesNames {
		fmt.Println(site.name)
	}
	s := mirroredQuery(sitesNames)
	fmt.Println("is:\n", s.name, "\nresponse is:", s.response)
}

// mirroredQuery makes goroutines for responses
func mirroredQuery(sites []siteMirror) siteMirror {
	responses := make(chan siteMirror, len(sites))

	for _, site := range sites {
		site := site // need new copy of site for every goroutine!
		go func() {
			responses <- request(site)
		}()
	}
	return <-responses // more fastest will return first
}

// request get site and calcs response
func request(site siteMirror) siteMirror {

	start := time.Now()
	//
	response, err := http.Get("http://" + site.name)
	check(err, "fatal", "Can't execute Get for "+site.name)
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	check(err, "fatal", "Can't ReadAll for "+site.name)
	//
	end := time.Now()
	site.response = end.Sub(start)

	return site
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
