// Homework-5: Standard library - Part 1
// Exercise 4 - MyCopy utility
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

var vFlag, yFlag, rFlag bool // -v -y -r
var fileCount int            // copied files counter

// MyCopy utility take two parameters: source and destination
// And some options: -y with no questions, -v verbose log, -r recursive coping folders
func main() {

	// set the custom Usage function
	flag.Usage = func() {
		fileName := filepath.Base(os.Args[0])
		fmt.Printf("usage: %s [options] <source> <destination>\n", fileName)
		fmt.Printf("example1: %s file1 file2\n", fileName)
		fmt.Printf("example2: %s -r dir1 dir2\n", fileName)
		flag.PrintDefaults()
	}

	flag.BoolVar(&yFlag, "y", false, "copy with no questions")
	flag.BoolVar(&vFlag, "v", false, "copy with verbose log")
	flag.BoolVar(&rFlag, "r", false, "recursive coping folders (without -r only single file can be copied)")
	flag.Parse()

	files := flag.Args() // source & dest files
	if len(os.Args) < 3 || len(files) != 2 {
		flag.Usage()
		os.Exit(2)
	}

	if !rFlag {
		copyFile(files[0], files[1]) // array is 2-elements only - src & dst
	} else {
		copyDir(files[0], files[1])
	}

	fmt.Println(fileCount, "files copied.")

}

// copyFile copies src file to dst file
func copyFile(src, dst string) {

	dstExists := true // let dst exists by default

	// open source file
	from, err := os.Open(src)
	check(err)
	defer from.Close()

	// check dst file exists
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		dstExists = false
	}

	// loop for getting overwrite agreement & check -y flag
	if !yFlag && dstExists {
		var a string
		for {
			fmt.Printf("File %s already exists, overwrite? [n|y]:", dst)
			fmt.Scanln(&a)
			a = strings.ToLower(a)
			if a == "n" {
				return
			} else if a == "y" {
				break // agree to overwrite
			}
		}
	}

	// create dst file
	to, err := os.Create(dst)
	check(err)
	defer to.Close()

	// copy routine
	count, err := io.Copy(to, from)
	check(err)

	// -v flag implement
	if vFlag {
		log.Printf(" - Copied %s -> %s - %d bytes", src, dst, count)
	}

	fileCount++
}

// copyDir copies src dir to dst dir
func copyDir(src, dst string) {

	// create dest dir include parents
	err := os.MkdirAll(dst, 0744) // do not move src to dst rights - for simplification
	check(err)

	dir, _ := os.Open(src)
	objects, err := dir.Readdir(-1) // single slice

	for _, file := range objects {

		srcPath := src + "/" + file.Name()
		dstPath := dst + "/" + file.Name()

		if file.IsDir() {
			// create sub-directories - recursively, sorry (
			copyDir(srcPath, dstPath)
		} else {
			// copy files
			copyFile(srcPath, dstPath)
		}
	}
}

// check simplifies the code when multiple treatments of the same type of errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
