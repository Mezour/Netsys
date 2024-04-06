package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	Dial, err := net.DialTimeout("tcp", "localhost:9009", 8*time.Second)
	if err != nil {
		panic(err)
	}

	defer Dial.Close()

	err = Dial.SetWriteDeadline(time.Now().Add(8 * time.Second))
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your message: ")
	scanner.Scan()
	payload := scanner.Text()

	_, err = Dial.Write([]byte(payload))
	if err != nil {
		panic(err)
	}

}
