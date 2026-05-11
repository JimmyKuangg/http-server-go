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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connecton: ", err)
			return
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected!")

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		message := string(buffer[:n])
		fmt.Println("Received: ", message)
	}
}