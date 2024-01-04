package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	reader := bufio.NewReader(os.Stdin)

	// Prompt for username
	fmt.Print("Enter username :")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Prompt for password
	fmt.Print("Enter password :")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Check username and password
	if strings.TrimSpace(username) == "std1" && strings.TrimSpace(password) == "p@ssw0rd" {
		fmt.Println("\n Server response : Hello \n")
	} else {
		fmt.Println("\n Server response : Invalid credentials \n")
		return
	}

	for {
		// Read user input
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Send the message to the server
		conn.Write([]byte(message))

		// Print the number of bytes sent
		fmt.Printf("Sent %d bytes\n", len(message))

		// Receive and print the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Server response: %s", buffer[:n])
	}
}
