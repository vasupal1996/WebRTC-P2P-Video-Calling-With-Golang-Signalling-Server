package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/pubsub"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func autoId() string {
	return uuid.NewV4().String()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var ps = &pubsub.PubSub{}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/", rootHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":8000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./static/html/index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true

	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go echo(conn)

}

func echo(conn *websocket.Conn) {
	client := pubsub.Client{
		Id:         autoId(),
		Connection: conn,
	}

	ps.AddClient(client)

	fmt.Println("New Client is connected, total: ", len(ps.Clients))
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Something went wrong", err)

			ps.RemoveClient(client)
			log.Println("total clients and subscriptions ", len(ps.Clients))

			return
		}

		ps.HandleReceiveMessage(client, messageType, p)

	}
}
