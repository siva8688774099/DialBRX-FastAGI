package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Siva_Nutakki/DialBRX-FastAGI/config"
	"github.com/Siva_Nutakki/DialBRX-FastAGI/internal/repositories"
)

func HandleAGIRequest(conn net.Conn) map[string]interface{} {
	defer conn.Close()

	// Placeholder for handling AGI requests
	// You would read from the connection, parse the AGI commands, and respond accordingly
	requestMapping := make(map[string]interface{})
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
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			requestMapping[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		fmt.Println("AGI ENV:", line)
	}
	fmt.Println("Final AGI ENV mapping:", requestMapping)

	_, err := conn.Write([]byte("200 result=0\n"))
	if err != nil {
		log.Println("Error writing AGI response:", err)
	}
	return requestMapping
}

func main() {
	config.LoadEnv()
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
		requestMapping := HandleAGIRequest(conn)
		fmt.Println("Received AGI request mapping:", requestMapping)
		// Do something with the request mapping, e.g., update call details
		if requestMapping["agi_network_script"] == "pushcallBackDetails" {
			err = repositories.PushCallBackDetails(requestMapping)
			if err != nil {
				log.Println("Error pushing call back details:", err)
			}
			return
		}
		if requestMapping["agi_network_script"] == "updatePostConnectDetails" {
			err = repositories.UpdateCallConnectDetails(requestMapping)
			if err != nil {
				log.Println("Error updating call connect details:", err)
			}
		}
	}
}
