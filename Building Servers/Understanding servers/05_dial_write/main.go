package main

import (
	"fmt"
	"net"
)

func main() {
	count := 0
	for {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(conn, "I dialed you")
		if count == 10 {
			break
		}
		count++
		defer conn.Close()
	}

}
