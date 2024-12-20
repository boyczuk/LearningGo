package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to existing server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	_, err = conn.Write([]byte(name))
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}
	fmt.Printf("Welcome, %s! Type and press enter to sent messages\n", name)

	// Go routine runs while waiting for user input
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error occured while reading from server: ", err)
				return
			}
			message := string(buffer[:n])
			fmt.Printf("\r%s\n> ", message)
		}
	}()

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		message := strings.TrimSpace(input)
		// Write to server conn
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}

	}
}
