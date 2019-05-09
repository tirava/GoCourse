// Homework-5: Standard library - Part 1
// Exercise 4 - Copy utility
// Author: Eugene Klimov
// Date: 10 may 2019
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var vFlag bool // -v
var yFlag bool // -y

// Copy utility take two parameters: source and destination
// And some options: -y with no questions, -v verbose log, -r recursive coping folders
func main() {

	// set the custom Usage function
	flag.Usage = func() {
		fileName := filepath.Base(os.Args[0])
		fmt.Printf("usage: %s [options] <source> <destination>\n", fileName)
		fmt.Printf("example1: %s file1 file2\n", fileName)
		//fmt.Printf("example2: %s -r dir1 dir2\n", fileName)
		flag.PrintDefaults()
	}

	flag.BoolVar(&yFlag, "y", false, "copy with no questions")
	flag.BoolVar(&vFlag, "v", false, "copy with verbose log")
	//rFlag := flag.Bool("r", false, "recursive coping folders (without -r only single file can be copied)")
	flag.Parse()

	files := flag.Args() // source & dest files
	if len(os.Args) < 3 || len(files) != 2 {
		flag.Usage()
		os.Exit(2)
	}

	copyFile(files[0], files[1]) // array is 2-elements only

	fmt.Println(yFlag, vFlag, flag.Args())

}

func copyFile(src, dst string) {

	dstExists := true // let dst exists by default

	from, err := os.Open(src)
	check(err)
	defer from.Close()

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		dstExists = false
	}

	// loop for getting overwrite agreement
	if !yFlag && dstExists {
		var a string
		for {
			fmt.Printf("File %s already exists, overwrite? [y|N]:", dst)
			fmt.Scanln(&a)
			a = strings.ToLower(a)
			if a == "n" {
				return
			} else if a == "y" {
				break // agree to overwrite
			}
		}
	}

	to, err := os.Create(dst)
	check(err)
	defer to.Close()

	count, err := io.Copy(to, from)
	check(err)

	if vFlag {
		log.Printf(" - Copied %s -> %s - %d bytes", src, dst, count)
	}
}

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
