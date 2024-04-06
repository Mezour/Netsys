package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:9009")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		connect, err := listener.Accept()
		err = connect.SetReadDeadline(time.Now().Add(8 * time.Second))
		if err != nil {
			panic(err)
		}

		go func(connection net.Conn) {
			fmt.Println("accepted")

			buff := make([]byte, 100)
			n, err := connection.Read(buff)
			if err != nil {
				panic(err)
			}
			fmt.Printf("read: %d, Message: %s\n", n, buff)
		}(connect)
	}

}
