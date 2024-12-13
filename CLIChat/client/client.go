package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		message := strings.TrimSpace(input)
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
		fmt.Println("Message sent to server.")
	}
}
