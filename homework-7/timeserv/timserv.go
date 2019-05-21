// Homework-7: Goroutines & Channels
// Exercise 3 - Time server
// Author: Eugene Klimov
// Date: 21 may 2019
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Time server with exit command
func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
		go handleRead(conn)
	}
}

// handleRead reads string from client with delimiter \n
func handleRead(c net.Conn) {
	defer c.Close()

	for {
		// read messages with \n
		msg, _ := bufio.NewReader(c).ReadString('\n')
		// convert to lower, cut \n and check
		msg = strings.ToLower(msg)
		msg = strings.TrimRight(msg, "\n")

		if msg == "exit" {
			// pong and exit
			_, err := c.Write([]byte("\nGood bye!\n"))
			if err != nil {
				log.Println(err)
			}
			c.Close()
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
