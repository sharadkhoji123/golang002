package main

import (
	"bufio"
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
		go Serve(conn)

	}

}

//Serve will serve the tcp connections
func Serve(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		fmt.Fprintf(conn, `Let me repeat what you said: "%s" \n  `, ln)
		if ln == "" {
			break
		}
	}

	defer conn.Close()

	fmt.Println("Code got here")
	io.WriteString(conn, "I see you connected")
}
