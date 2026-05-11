package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	requestLine = strings.TrimSpace(requestLine)

	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request line")
	}

	method := parts[0]
	path := parts[1]

	// headers
	headers := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		hparts := strings.SplitN(line, ":", 2)
		if len(hparts) != 2 {
			continue
		}

		headers[strings.TrimSpace(hparts[0])] = strings.TrimSpace(hparts[1])
	}

	// body
	contentLength := 0
	if val, ok := headers["Content-Length"]; ok {
		fmt.Sscanf(val, "%d", &contentLength)
	}

	body := make([]byte, contentLength)

	if contentLength > 0 {
		_, err := reader.Read(body)
		if err != nil {
			return nil, err
		}
	}

	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    body,
	}, nil
}