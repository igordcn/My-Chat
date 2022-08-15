package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var connexions = make(map[string]*websocket.Conn)
var count = ""

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/message", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	count = count + "1"
	id := count
	connexions[id] = conn
	defer delete(connexions, id)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		for _, value := range connexions {
			if value != conn {
				err2 := value.WriteMessage(websocket.TextMessage, message)
				if err2 != nil {
					log.Println(err2)
					return
				}
			}
		}
	}
}
