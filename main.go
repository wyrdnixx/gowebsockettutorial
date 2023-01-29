package main

// https://www.youtube.com/watch?v=norUcMSJRtQ

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting app")

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	ws, err := NewWebSocket(w, r)
	// 	if err != nil {
	// 		log.Printf("Error creating websocket connection: %v", err)
	// 	}
	// 	ws.On("message", func(e *Event) {
	// 		log.Printf("Message received: %s", e.Data.(string))

	// 		// create event and just send the message back
	// 		ws.Out <- (&Event{
	// 			Name: "response",
	// 			//Data: e.Data.(string)
	// 			Data: strings.ToUpper(e.Data.(string)),
	// 		}).Raw()
	// 	})

	// })

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error creating websocket connection: %v", err)
		}
		webSocketEventHandler(*ws)
	})

	http.ListenAndServe(":8080", nil)

}
