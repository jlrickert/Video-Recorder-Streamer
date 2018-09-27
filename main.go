package main

import (
	"log"
	"net/http"
)

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func main() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("connecting")
		serveWs(hub, w, r)
	})
	// err := http.ListenAndServe(":8090", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	err := http.ListenAndServeTLS("localhost:8090", "video_streamer.crt", "video_streamer.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
