package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err.Error())
		}
		go handConnection(conn)
	}
}

func handConnection(conn net.Conn) {
	// Echo all incoming data.
	io.Copy(conn, conn)
	// Shut down the connection.
	conn.Close()
	// var b []byte
	// for {
	// err := conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	// if err == nil {
	// 	fmt.Print(err)
	// }
	// n, err := conn.Read(b)
	// fmt.Print(n, err)
	// if n > 10 {
	// 	break
	// }
	// }

	// conn.Close()
}
