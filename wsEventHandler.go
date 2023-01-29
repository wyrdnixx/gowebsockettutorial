package main

import (
	"log"
	"strings"
)

func webSocketEventHandler(ws WebSocket) {

	ws.On("TestMessageReply", func(e *Event) {
		log.Printf("Message received: %s", e.Data.(string))

		// create event and just send the message back
		ws.Out <- (&Event{
			Name: "response",
			//Data: e.Data.(string)
			Data: strings.ToUpper(e.Data.(string)),
		}).Raw()
	})

	// for ev := range ws.Events {
	// 	log.Printf("event;: %v ", ev)
	// }

}
