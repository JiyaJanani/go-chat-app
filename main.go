package main

import (
	"flag"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		//index 0 is having the name of the program 1 is teh listen flag 2 is the ip
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("error:", listenErr)
	}

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("error:", acceptErr)
	}

}

func runGuest(ip string) {

}
