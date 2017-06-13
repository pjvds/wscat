package main

import (
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var origin = "http://localhost/"

func main() {
	if len(os.Args) == 1 {
		log.Fatal("missing url argument, use: wscat ws://localhost")
	}

	url := os.Args[1]
	if !strings.HasPrefix("ws://", url) {
		url = "ws://" + url
	}

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("%s\n", message)
	}
}
