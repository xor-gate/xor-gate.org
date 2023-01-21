package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// The import statement is used to import packages
//  imports from third party sources are always URLs
//  so you know where they are from
//---
// EXERCISE 1: must import the packages which are used throughout the code
//---

var hostPort = "127.0.0.1:9001" // IoT device TCP server address
const clientsToStart = 16       // Amount of fake clients to spawn
const clientStartWaitTime = 1 * time.Millisecond

// The main function where a little magic happens
func main() {
	var clientID uint = 1
	var allClientsStarted bool

	clientsDisconnectChannel := make(chan string, clientsToStart)

	// Run the server on its own goroutine (thread)
	//---
	// EXERCISE 2: Run the TCP server "in the background" with a goroutine
	//---
	go server_listen_and_serve()

	// Because the server is started asynchronous we don't know when we are
	//  ready to accept client connections. For now we sleep a little. In
	//  real life the server is started before the clients connect
	time.Sleep(3 * time.Second)

	// Start the amount of clientsToStart
	for {
		//---
		// EXERCISE 2: Run the TCP client "in the background" with a goroutine
		//---
		go client(clientID, clientsDisconnectChannel) // spawn client
		time.Sleep(clientStartWaitTime)               // sleep fixed time for spawning next client

		// Generate next client ID
		clientID++
		if clientID > clientsToStart {
			log.Println("All clients are started in parallel!")
			allClientsStarted = true
			break
		}
	}

	// When all clients are spawned in parallel wait for them to get the
	//  disconnect message from the server
	if allClientsStarted {
		var clientsDisconnected = 0     // Count the amount of clients disconnected
		var allClientsDisconnected bool // We are done when all clients signaled and disconnected

		for {
			select {
			case <-time.After(5 * time.Second):
				panic("Clients not disconnected within 5 seconds!")
			case clientDisconnectMsg := <-clientsDisconnectChannel:
				log.Println(clientDisconnectMsg)
				clientsDisconnected++
				if clientsDisconnected == clientsToStart-1 {
					allClientsDisconnected = true
					break
				}
			}

			if allClientsDisconnected {
				break
			}
		}
	} else {
		log.Println("Client sync lost")
	}

	log.Println("All", clientsToStart, "clients disconnected")
}

// TCP client implementation
// clientID the client ID when connecting to the server
func client(clientID uint, disconnectChannel chan string) {
	conn, err := net.Dial("tcp", hostPort)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	// Succesful connect to server, send our id as string: "id=<clientID\n"
	log.Printf(fmt.Sprintf("[client] Started client with ID %d\n", clientID))
	conn.Write([]byte(fmt.Sprintf("id=%v\n", clientID)))

	// Read response from server
	response := bufio.NewReader(conn)
	for {
		switch err {
		case nil:
			conn.Write([]byte(fmt.Sprintf("[client] id=%u", clientID)))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("ERROR", err)

		}

		responseData, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			responseString := string(responseData)
			switch responseString {
			case "DISCONNECT":
				// Notify the main goroutine the client connection was ok, and is disconnected
				disconnectChannel <- fmt.Sprintf("[client] id=%d got \"", responseString, "\" msg from server", clientID)
				conn.Close() // Close the client side connection
			}
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("ERROR", err)
			os.Exit(2)
		}
	}
}

// This is the command and control server for a single client
//
//	the conn parameter is the connection to the client
//
// As this function is blocking it must be executed from a goroutine
func serve_single_client(conn net.Conn) {
	log.Println("[server] TCP client connected")

	r := bufio.NewReader(conn)
	for {
		// Receive a "hello" message from the client (every message ends with a newline)
		data, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("ERROR", err) // Something went wrong while reading from the server e.g disconnect
			return
		}

		msg := string(data)
		log.Println("[server] Received client message:", msg)

		// Ask the client to disconnect
		//conn.Write([]byte("DISCONNECT"))
	}

	// Wait some time to flush the connection buffer and close the client TCP connection
	time.Sleep(time.Second)
	conn.Close()
}

func server_listen_and_serve() {
	// Start a TCP server
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		fmt.Println("[server] ERROR while started to listen on", hostPort, "error:", err)
		os.Exit(1)
	}

	fmt.Println("Waiting for TCP clients at", hostPort)

	for {
		conn, err := l.Accept() // Wait for connection of a client
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		// Serve a single client on its own goroutine
		go serve_single_client(conn)
	}
}

