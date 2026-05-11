package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting TCP server...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}

	fmt.Println("Listening on :8080")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connecton: ", err)
		return
	}

	fmt.Println("Client connected!")

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from buffer: ", err)
		return
	}

	fmt.Println("Received bytes:")
	fmt.Println(string(buffer[:n]))
}