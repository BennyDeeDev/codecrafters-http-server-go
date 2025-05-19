package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func HandleConnection(c net.Conn) {
	defer c.Close()

	req := make([]byte, 1024)

	c.Read(req)

	fmt.Println(string(req))

	s := strings.Fields(string(req))
	fmt.Println(s)

	var res string

	p := strings.Split(s[1], "/")

	if p[1] == "echo" {
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(p[2]), p[2])
	} else if p[1] == "user-agent" {
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(s[6]), s[6])
	} else if s[1] != "/" {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	} else {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	}

	c.Write([]byte(res))
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Running HTTP Server on 4221")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleConnection(c)
	}
}
