package main

import (
	"fmt"
	"net"
	"strings"
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

		input := strings.TrimSpace(string(buffer[:n]))
		
		parts := strings.Split(input, " ")

		if len(parts) != 2 {
    	conn.Write([]byte("bad request format\n"))
    	continue
		}

		method := parts[0]
		path := parts[1]

		fmt.Printf("Method: %s, Path: %s\n", method, path)

		var response string

		if method != "GET" {
			response = "only GET supported\n"
		} else {
			switch path {
			case "/ping":
				response = "PONG\n"
			case "hello":
				response = "hello human 👋\n"
			default:
				response = "unknown command\n"
			}
		}

		_, err = conn.Write([]byte(response)) 
		if err != nil {
			fmt.Println("Error writing back to client: ", err)
			return
		}
	}
}

func trim(s string) string {
    // remove \n and \r from nc input
    return strings.TrimSpace(s)
}