package main

import "encoding/json"

type EventHandler func(*Event)

/*
EXAMPLE
{
	"event": "message"
	"data": "THis ist a data object"
}

*/
type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
