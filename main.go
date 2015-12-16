package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"

func main() {
	if len(os.Args) == 0 {
		log.Fatal("missing url argument, use: wscat ws://localhost")
	}

	socket, err := websocket.Dial(os.Args[1], "", origin)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1000)
	for {
		n, err := socket.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buffer[0:n])
	}
}
