package main

import (
	"fmt"
	"net"
)

func main() {
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
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client or client disconnected:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println(message)
	}
}
