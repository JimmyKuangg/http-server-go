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

	router := NewRouter()

	router.GET("/ping", func(req *Request) string {
		return string(req.Body)
	})

	router.GET("/hello", func(req *Request) string {
		return "hello from server 👋"
	})

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connecton: ", err)
			return
		}

		go handleConnection(conn, router)
	}
}

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()


	req, err := parseRequest(conn)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	responseBody := router.Handle(req)

	writeResponse(conn, "200 OK", responseBody)
}