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

const fileName = "mirrored.txt"

var sitesNames []string
var respTime time.Duration

// Reads sites from txt file and checks more fastest
func main() {

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
		sitesNames = append(sitesNames, line)
	}

	fmt.Println("Fastest site/mirror from:\n", sitesNames, "is:\n", mirroredQuery(sitesNames), "\nresponse is:", respTime)
	// +ms
}

func mirroredQuery(sites []string) string {
	responses := make(chan string, len(sites))

	for _, site := range sites {
		site := strings.TrimRight(site, "\n") // need new copy of site name for every goroutine!
		go func() {
			responses <- request(site)
		}()
	}
	return <-responses // more fastest will return first
}

func request(hostname string) string {

	start := time.Now()
	//
	response, err := http.Get("http://" + hostname)
	check(err, "fatal", "Can't execute Get for "+hostname)
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	check(err, "fatal", "Can't ReadAll for "+hostname)
	//
	end := time.Now()
	respTime = end.Sub(start)

	return hostname
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
