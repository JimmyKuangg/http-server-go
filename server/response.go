package main

import (
	"fmt"
	"net"
)

type Response struct {
	Status int
	Body   any
}

func writeResponse(conn net.Conn, res Response) {
	bodyStr := fmt.Sprintf("%v", res.Body)
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", res.Status, statusText(res.Status))

	response :=
		statusLine +
			"Content-Type: text/plain\r\n" +
			"Content-Length: " + fmt.Sprint(len(bodyStr)) + "\r\n" +
			"\r\n" +
			bodyStr

	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error writing to conn: %v", err)
	}
}

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}