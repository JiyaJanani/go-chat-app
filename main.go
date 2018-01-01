package main

import (
	"Z:\go\chat-app\lib\main.go"
	"flag"
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
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
