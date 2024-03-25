package main

import (
	"fmt"
	"net"
)

func main() {

	Dial, err := net.Dial("tcp", "localhost:9999")

	// error handling wajib
	if err != nil {
		return
	}

	defer Dial.Close()

	payload := "Hello World\n"

	_, err = Dial.Write([]byte(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

}
