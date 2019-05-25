// Homework-8: Goroutines & Channels - practice
// Exercise 1 - Serial sites scanner
// Author: Eugene Klimov
// Date: 25 may 2019
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 2 { // no args
		fmt.Println("Run " + filepath.Base(os.Args[0]) + " <file with sites txt>")
		os.Exit(2)
	}

	sites := readSites(os.Args[1])

	start := time.Now()
	for _, url := range sites {
		fetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get("http://" + url)
	check(err, "fatal", "Can't execute Get for "+url)

	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	err = resp.Body.Close()
	check(err, "fatal", "Error body closing!")

	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

// readSites reads sites from file and returns slice
func readSites(fileName string) []string {
	sitesNames := make([]string, 0)

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
		if len(line) < 3 { // no fake symbols
			continue
		}
		sitesNames = append(sitesNames, strings.TrimRight(line, "\n"))
	}

	return sitesNames
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
