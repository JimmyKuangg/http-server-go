package main

import (
	"fmt"
	"net"
)

func writeResponse(conn net.Conn, status string, body string) {
	response :=
		"HTTP/1.1 " + status + "\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: " + fmt.Sprint(len(body)) + "\r\n" +
			"\r\n" +
			body

	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error writing to conn: %v", err)
	}
}