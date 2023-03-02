package main

import (
	"encoding/json"
	"log"
	"strings"
)

func webSocketEventHandler(ws WebSocket) {

	ws.On("TestMessageReply", func(e *Event) {
		log.Printf("Message received: %s", e.Data.(string))
		//log.Printf("test")
		// create event and just send the message back
		ws.Out <- (&Event{
			Name: "response",
			//Data: e.Data.(string)
			Data: strings.ToUpper(e.Data.(string)),
		}).Raw()
	})

	ws.On("SimpleOkResponse", func(e *Event) {
		raw, _ := json.Marshal(e.Data)
		log.Printf("%s", raw)
		ws.Out <- (&Event{
			Name: "response",
			Data: "ok",
		}).Raw()
	})

	// for ev := range ws.Events {
	// 	log.Printf("event;: %v ", ev)
	// }

}
