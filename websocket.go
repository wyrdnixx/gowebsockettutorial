package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {

		//ToDo : Checks bauen - Header oder origin etc...
		return true
	},
}

type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

type errorMessage struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error while upgrading connection: %v", err)
		return nil, err
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	go ws.Reader()
	go ws.Writer()
	return ws, nil
}

func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()

	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error while reading data from websocket - closed: %v", err)
			} else {
				log.Printf("Error while reading data from websocket - other error: %v", err)
			}
			break
		}
		event, err := NewEventFromRaw((message))
		if err != nil {
			log.Printf("Error parsing message: %v", err)
		} else {
			//log.Printf("got message: %v", message)
		}
		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		} else {
			log.Printf("Websocket reader received unknown event or message from client: %v", bytes.NewBuffer(message).String())

			x := new(errorMessage)
			x.Name = "error"
			x.Error = "Websocket reader received unknown event or message from client"

			//ws.Conn.WriteJSON(x)
			raw, _ := json.Marshal(x)
			ws.Out <- raw

		}
	}
}

func (ws *WebSocket) Writer() {
	for {
		select {
		case message, ok := <-ws.Out:
			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				return
			}
			w, err := ws.Conn.NextWriter((websocket.TextMessage))
			if err != nil {
				return
			}
			w.Write((message))
			w.Close()

		}
	}
}

func (ws *WebSocket) On(eventName string, action EventHandler) *WebSocket {
	ws.Events[eventName] = action
	return ws
}
