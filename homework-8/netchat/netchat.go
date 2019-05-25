// Homework-8: Goroutines & Channels - practice
// Exercise 4 - Net chat modifications
// Author: Eugene Klimov
// Date: 25 may 2019
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	hostPort    = "localhost:8000"
	exitSeconds = 5
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// Added from netchat.txt:
// 1. Server run Hello string
// 2. Normal exiting server
// 3. Send client notifications on server exit
// 4. Log clint connect/disconnect
// 5. User friendly client names
// 6. Client exiting by 'exit'
func main() {
	listener, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatal(err)
	}
	go waitExit()

	log.Println("Hello! I'm running on:", hostPort)

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	whoAddr := conn.RemoteAddr().String()
	who := getUserName(whoAddr)
	log.Println("Connected -", whoAddr, "as", who)
	ch <- "You are " + who + "\nType 'exit' for disconnect"
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := input.Text()
		if strings.ToLower(s) == "exit" {
			ch <- "Good bye!"
			time.Sleep(time.Second) // need wait for the message printing before leaving (waitgroup for writing)
			break
		}
		messages <- who + ": " + s
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
	log.Println("Disconnected -", whoAddr, "as", who)
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func waitExit() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == "exit" {
			log.Println("Sending to clients messages about", exitSeconds, "seconds...")
			messages <- "Server shutdown in " + strconv.Itoa(exitSeconds) + " seconds, Sorry..."
			time.Sleep(exitSeconds * time.Second)
			log.Println("Exiting... Good bye!")
			os.Exit(0)
		}
	}
}

// getUserName returns random user name
func getUserName(addr string) (user string) {

	resp, err := http.Get("http://randomuser.ru/api.json")
	check(err, "", "Error get users url!")
	bytes, err := ioutil.ReadAll(resp.Body)
	check(err, "", "Error read users body!")

	if err != nil {
		return addr
	}

	draft := string(bytes)
	s := strings.Split(draft, ":")  // draft names
	sF := strings.Split(s[4], "\"") // first
	sL := strings.Split(s[5], "\"") // last
	user = sF[1] + " " + sL[1]
	return
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
