package main

import (
	"fmt"
	"net"
	"sync"
)

var clients = make(map[net.Conn]string)
var mutex = &sync.Mutex{}

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

		// Go routine
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	// Remove clients from client map and close connection on function end
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	// Create buffer to store input
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading user's name:", err)
		return
	}
	name := string(buffer[:n])
	fmt.Printf("New client connected: %s\n", name)

	mutex.Lock()
	clients[conn] = name
	mutex.Unlock()

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

		// Convert to message, send name of connected client
		broadcastMessage(fmt.Sprintf("%s: %s", name, message), conn)
	}
}

func broadcastMessage(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for conn := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending message to client:", err)
			}
		}
	}
}
