package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

//RunHost takes an ip as argument
//and listen for connection on that ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("error:", listenErr)
	}

	fmt.Println("Listening on", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("error:", acceptErr)
	}
	fmt.Println("new connection Accepted")
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message Received ", message)

}

// RunGuest takes destination ip
//as argument and connects to that ip
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dailErr != nil {
		log.Fatal("Error: ", dailErr)
	}
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr = reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint()
}
