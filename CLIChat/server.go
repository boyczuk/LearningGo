package main

import (
	"fmt"
	"net"
)

func main() {
	// Create server
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		// defer conn.Close()
		fmt.Println("Client connected!")

		// Go routine
		go handleClient(conn)

		// buffer := make([]byte, 1024)
		// n, err := conn.Read(buffer)
		// if err != nil {
		// 	fmt.Println("Error reading from client:", err)
		// 	return
		// }

		// message := string(buffer[:n])
		// fmt.Println(message)
	}

}

func handleClient(conn net.Conn) {
	// Close connections once done
	defer conn.Close()
	// Create buffer to store input
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading user's name:", err)
		return
	}
	name := string(buffer[:n])
	fmt.Printf("New client connected: %s\n", name)

	for {
		// Read the information sent from the client into the buffer
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("User %s disconnected:\n", name)
			return
		}

		// convert slice buffer into string and print
		message := string(buffer[:n])
		fmt.Printf("%s: %s\n", name, message)
	}
}
