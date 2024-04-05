package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	Dial, err := net.DialTimeout("tcp", "localhost:9009", 5*time.Second)
	// error handling wajib
	if err != nil {
		return
	}

	defer Dial.Close()

	err = Dial.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return
	}

	err = Dial.SetWriteDeadline(time.Now().Add(8 * time.Second))
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your message: ")
	scanner.Scan()
	payload := scanner.Text()

	// payload := "Hello World\n"

	_, err = Dial.Write([]byte(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

}
