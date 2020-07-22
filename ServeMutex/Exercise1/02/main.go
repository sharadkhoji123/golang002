package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
	io.WriteString(conn, "\nI see you connected\n")
	fmt.Fprintln(conn, "How is your day?")
	fmt.Fprintf(conn, "Well I hope!")

	conn.Close()
}
