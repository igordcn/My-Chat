package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

var connexions = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connexions["usuario1"] = conn
	defer conn.Close()
	defer delete(connexions, "usuario1")

	cw, err2 := conn.NextWriter(websocket.TextMessage)
	defer cw.Close()
	if err2 != nil {
		log.Println(err2)
		return
	}

	cw.Write("")

}
