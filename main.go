package main

// https://www.youtube.com/watch?v=norUcMSJRtQ

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting app")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error creating websocket connection: %v", err)
		}
		ws.On("message", func(e *Event) {
			log.Printf("Message received: %s", e.Data.(string))
		})

		ws.On("message2", func(e *Event) {
			log.Printf("Message2 received: %s", e.Data.(string))
		})
	})

	http.ListenAndServe(":8080", nil)

}
