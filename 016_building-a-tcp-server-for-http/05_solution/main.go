package main

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URL", u)

	// multiplexer
	if m == "GET" {
		switch u {
		case "/":
			index(conn)
		case "/foo":
			foo(conn)
		default:
			notFound(conn)
		}
	} else if m == "POST" {
		switch u {
		case "/foo":
			fooProcess(conn)
		}
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Index</title>
</head>
<body>
<strong>Index</strong>
<ul>
	<li><a href="/">Index</a></li>
	<li><a href="/foo">Foo</a></li>
	<li><a href="/doesNotExist">Something here?</a></li>
</ul>
</body>
</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func foo(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Are you a foo?</title>
</head>
<body>
<strong>Are you a foo?</strong>
<ul>
	<li><a href="/">Index</a></li>
	<li><a href="/foo">Foo</a></li>
	<li><a href="/doesNotExist">Something here?</a></li>
</ul>
<form method="POST" action="/foo">
<input type="Submit" value="Yes">
</form>
</body>
</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func fooProcess(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Are you a foo?</title>
</head>
<body>
<strong>Are you a foo?</strong>
<ul>
	<li><a href="/">Index</a></li>
	<li><a href="/foo">Foo</a></li>
	<li><a href="/doesNotExist">Something here?</a></li>
</ul>
<p>
	Congratulations! :)
</p>
</body>
</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func notFound(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>404 Not found</title>
</head>
<body>
<strong>404 Not found</strong>
<ul>
	<li><a href="/">Index</a></li>
	<li><a href="/foo">Foo</a></li>
	<li><a href="/doesNotExist">Something here?</a></li>
</ul>
</body>
</html>`

	fmt.Fprint(conn, "HTTP/1.1 404 Not found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
