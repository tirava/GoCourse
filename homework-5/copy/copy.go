// Homework-5: Standard library - Part 1
// Exercise 4 - Copy utility
// Author: Eugene Klimov
// Date: 10 may 2019
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Copy utility take two parameters: source and destination
// And some options: -y with no questions, -v verbose log, -r recursive coping folders
func main() {

	yFlag := flag.Bool("y", false, "copy with no questions")
	vFlag := flag.Bool("v", false, "copy with verbose log")
	rFlag := flag.Bool("r", false, "recursive coping folders")
	flag.Parse()

	files := flag.Args() // source & dest files
	if len(os.Args) < 3 || len(files) != 2 {
		usage()
	}

	fmt.Println(*yFlag, *vFlag, *rFlag, flag.Args())

}

// usage overrides default helper
var usage = func() {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fmt.Fprintf(cmdLine.Output(), "usage: %s [options] <source> <destination>\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}
