package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			// When ln is empty, header is done
			fmt.Println("End of HTTP request headers")
			break
		}
		i++
	}

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Low level html with Go</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
