package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Buffer for reading
	buffer := make([]byte, 1024)

	// Read username and password from the client
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading credentials:", err)
		return
	}

	// Extract username and password from the received data
	credentials := strings.Split(string(buffer[:n]), ":")
	if len(credentials) != 2 {
		fmt.Println("Invalid credentials format")
		conn.Write([]byte("Invalid credentials\n"))
		return
	}

	// Check username and password
	username := strings.TrimSpace(credentials[0])
	password := strings.TrimSpace(credentials[1])

	if username == "std1" && password == "p@ssw0rd" {
		// Valid credentials, respond with "Hello"
		response := "Hello\n"
		conn.Write([]byte(response))
	} else {
		// Invalid credentials, respond with "Invalid credentials"
		response := "Invalid credentials\n"
		conn.Write([]byte(response))
	}
}

func main() {
	//Listen for incoming connections
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	//Close the listenner when the application
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	//Listen for an incoming connection

	for {

		//Accept() block until a connection is made
		conn, err := listener.Accept() //Accept it three-way handshark
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue //continue to next
		}

		fmt.Println("New connection established")

		go handleConnection(conn)
	}
}
