package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	w, err := conn.NextWriter()
	if err != nil {
		log.Println(err)
		return
	}

}
