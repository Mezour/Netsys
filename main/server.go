package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:9999")

	// error handling wajib
	if err != nil {
		return
	}

	defer listener.Close()

	for {
		connect, err := listener.Accept()

		if err != nil {
			return
		}

		go func(connection net.Conn) {
			fmt.Println("accepted")

			buff := make([]byte, 100)

			//read payload that has been sent
			//depan read dibawah pakai nama variable

			n, err := connection.Read(buff)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("read: %d, Message: %s\n", n, buff)

		}(connect)
	}

}
