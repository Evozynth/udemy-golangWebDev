package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"strings"
	"fmt"
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
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN_MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDELETE key \r\n"+
		"EXAMPLE:\r\n"+
		"SET fav chocolate \r\n"+
		"GET fav \r\n\r\n\r\n")

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// logic
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			if len(fs) == 1 {
				for k, v := range data {
					fmt.Fprintf(conn, "%s - %s\r\n", k, v)
				}
			} else {
				k := fs[1]
				v := data[k]
				fmt.Fprintf(conn, "%s\r\n", v)
			}
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXCPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
