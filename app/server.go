package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, 1024)

	_, err := c.Read(data)
	if err != nil {
		fmt.Println("Error reading request connection", err.Error())
		return
	}

	pong := formatRESPString("PONG")
	_, err = c.Write([]byte(pong))
	if err != nil {
		fmt.Println("Error writing respnose", err.Error())
		return
	}
}

func formatRESPString(str string) string {
	return fmt.Sprintf("+%s\r\n", str)
}
