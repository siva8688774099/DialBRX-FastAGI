package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleAGIRequest(conn net.Conn) {
	// Placeholder for handling AGI requests
	// You would read from the connection, parse the AGI commands, and respond accordingly
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading line:", err)
			break
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		fmt.Println("AGI ENV:", line)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":4573")
	if err != nil {
		log.Fatal("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("FastAGI server running on port 4573\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}
		go HandleAGIRequest(conn)
	}
}
