// Homework-8: Goroutines & Channels - practice
// Exercise 1 - Parallel sites scanner
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
	ch := make(chan string)
	for _, url := range sites {
		go fetch(url, ch)
	}
	for range sites {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get("http://" + url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

// readSites reads sites from file and returns slice
func readSites(fileName string) []string {
	sitesNames := make([]string, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
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
