package main

import (
	"bufio"
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request line: ", err)
		return
	}
	requestLine = strings.TrimSpace(requestLine)

	headers := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading headers: %s", err)
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		headers[key] = value
	}

	fmt.Println("Headers:", headers)

	parts := strings.Split(requestLine, " ")

	if len(parts) < 3 {
			fmt.Println("bad request")
			return
	}

	method := parts[0]
	path := parts[1]


	if method != "GET" {
			writeResponse(conn, "405 Method Not Allowed", "only GET supported")
			return
	}

	switch path {
	case "/ping":
			writeResponse(conn, "200 OK", "pong")
	case "/hello":
			writeResponse(conn, "200 OK", "hello from server 👋")
	default:
			writeResponse(conn, "404 Not Found", "not found")
	}
}

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