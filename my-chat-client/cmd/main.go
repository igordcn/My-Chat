package main

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/message"}
	log.Printf("Connecting to %s\n", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Server Does Not Found Error: ", err)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Message Receiving Error: ", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		text, _ := reader.ReadString('\n')
		err := c.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("Message Sending Error: ", err)
		}
	}

}
