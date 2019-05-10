// Homework-5: Standard library - Part 1
// Exercise 4 - MyGrep utility
// Author: Eugene Klimov
// Автор: Евгений Климов (строка для тестов grep)
// Date: 11 may 2019
package main

import (
	"bufio"
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
		fmt.Printf("example1: dir(ls) | %s my\n", fileName)
		fmt.Printf("example2: %s -i klim mygrep.go\n", fileName)
		flag.PrintDefaults()
	}

	flag.BoolVar(&iFlag, "i", false, "search case insensitive")
	flag.Parse()

	searchStr := flag.Args() // search string - array
	l := len(searchStr)
	if len(os.Args) < 2 || l < 1 || l > 2 {
		flag.Usage()
		os.Exit(2)
	}

	// check file name is present and search in it
	if l == 2 {
		bs, err := ioutil.ReadFile(searchStr[1])
		check(err)
		fmt.Println(findAllStrings(string(bs), searchStr[0]))
	}

	// search from stdin
	if l == 1 {
		fmt.Println("Not implemented", searchStr[0])
	}
}

// findAllStrings finds all strings by search template and returns these
func findAllStrings(inStr, searchStr string) string {
	var outLines string

	// get lines slice from clear string
	lines := stringToLines(&inStr)

	// get next line
	// check line for search string
	// if found - append to out string
	for _, line := range lines {
		if strings.Contains(line, searchStr) {
			//outLines = append(outLines, line)
			outLines = outLines + line + "\n"
		}
	}

	return outLines
}

// stringToLines convert string to lines slice
func stringToLines(str *string) (lines []string) {

	scanner := bufio.NewScanner(strings.NewReader(*str))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return
}

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
