package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
	}()
	_, _ = io.Copy(conn, os.Stdin) // until you send ^Z - does not work in Linux!
	fmt.Printf("%s: exited\n", conn.LocalAddr())
}
