package main

import (
	"net"
	"log"
	"io"
	"bufio"
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
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				// When ln is empty, header is done
				fmt.Println("End of HTTP request headers")
				break
			}
		}

		fmt.Println("Code got here")
		io.WriteString(conn, "\nI see you connected\n")

		conn.Close()
	}
}
