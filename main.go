package main

// https://www.youtube.com/watch?v=norUcMSJRtQ

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {

	log.Println("Starting app")

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error creating websocket connection: %v", err)
		}
		ws.On("message", func(e *Event) {
			log.Printf("Message received: %s", e.Data.(string))

			// create event and just send the message back
			ws.Out <- (&Event{
				Name: "response",
				//Data: e.Data.(string)
				Data: strings.ToUpper(e.Data.(string)),
			}).Raw()
		})

		ws.On("message2", func(e *Event) {
			log.Printf("Message2 received: %s", e.Data.(string))

			// own testanswer - return the string got from event

			type Answer struct {
				Name string `json:"name"`
				//Data map[string]interface{}
				Data map[string]interface{} `json:"data"`
			}

			aw := new(Answer)
			aw.Name = "reply"

			type foo struct {
				X     string `json:"X"`
				Y     string `json:"Y"`
				Z     string `json:"Z"`
				Reply string `json:"reply"`
			}

			bar := new(foo)
			bar.Reply = e.Data.(string)
			bar.X = "X"
			bar.Y = "Y"

			// dat := map[string]interface{}{

			// 	"X":     "a",
			// 	"Y":     "b",
			// 	"Z":     "c",
			// 	"reply": e.Data.(string),
			// }

			//warum geht das so, aber direkt dt in Answer.Data nicht?
			dat2 := map[string]interface{}{
				"values": bar,
			}

			aw.Data = dat2
			//aw.Data = map[string]interface{}{"data": bar}
			//dat2

			jAW, _ := json.Marshal(aw)

			ws.Out <- jAW

		})
	})

	http.ListenAndServe(":8080", nil)

}
