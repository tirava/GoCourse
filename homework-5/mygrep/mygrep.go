// Homework-5: Standard library - Part 1
// Exercise 4 - MyGrep utility
// Author: Eugene Klimov
// Автор: Евгений Климов (строка для тестов grep)
// Date: 11 may 2019
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var iFlag bool

// MyGrep utility is simple analog of world-wide known linux grep
func main() {

	// set the custom Usage function
	flag.Usage = func() {
		fileName := filepath.Base(os.Args[0])
		fmt.Printf("usage1: <command> | %s [options] <search string>\n", fileName)
		fmt.Printf("usage2: %s [options] <search string> <file>\n", fileName)
		fmt.Printf("example1: dir | %s my\n", fileName)
		fmt.Printf("example2: cat maygrep.go | ./%s Klim\n", fileName)
		fmt.Printf("example3: %s -i klim mygrep.go\n", fileName)
		flag.PrintDefaults()
	}

	flag.BoolVar(&iFlag, "i", false, "search case insensitive")
	flag.Parse()

	searchStr := flag.Args() // array with search string or +file name
	l := len(searchStr)
	if len(os.Args) < 2 || l < 1 || l > 2 {
		flag.Usage()
		os.Exit(2)
	}

	// check file name is present and search in it
	var inString string
	if l == 2 {
		bs, err := ioutil.ReadFile(searchStr[1])
		check(err)
		inString = string(bs)
	}

	// search from stdin
	if l == 1 {
		fi, _ := os.Stdin.Stat()
		// check if stdin from pipe
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			bytes, err := ioutil.ReadAll(os.Stdin)
			check(err)
			inString = string(bytes)
		} else { // from terminal - exit
			flag.Usage()
			os.Exit(2)
		}
	}

	fmt.Println(findAllStrings(&inString, searchStr[0]))
}

// findAllStrings finds all strings by search template and returns these
func findAllStrings(inStr *string, searchStr string) string {
	var outLines, l string

	// check -i options
	if iFlag {
		searchStr = strings.ToLower(searchStr)
	}

	// get lines slice from clear string
	lines := strings.Split(*inStr, "\n")

	// loop checks every line for search string
	for _, line := range lines {
		// check -i options
		if iFlag {
			l = strings.ToLower(line)
		} else {
			l = line
		}
		// check l but append line
		if strings.Contains(l, searchStr) {
			outLines = outLines + line + "\n"
		}
	}

	return outLines
}

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
