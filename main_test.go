package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/websocket"
)

var addrCLI = flag.String("addrCLI", "localhost:8080", "http service address")
var ws *websocket.Conn

func TestMain(m *testing.M) {
	fmt.Println("runiing testmain...")
	go main()

	u := url.URL{Scheme: "ws", Host: *addrCLI, Path: "/ws"}
	wsInit, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)

	}
	ws = wsInit

	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func TestWsEchoMessage(t *testing.T) {
	testMessage := Event{
		Name: "TestMessageReply",
		Data: `this is test data`,
	}
	err := ws.WriteMessage(1, testMessage.Raw())

	if err != nil {
		//t.Fatalf(`Sendmessage got : %v,  "", error`, err)
		t.Fatalf(`Sendmessage got error : %v  `, err)

	} else {
		t.Log("testWS succeded")
	}

	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("error reading reply message: %v", err)
	} else {
		expextedReply := `{"event":"response","data":"THIS IS TEST DATA"}`
		if string(message) != expextedReply {
			t.Fatalf("expected response : %v , got isntead: %v", expextedReply, string(message))
		}

		t.Logf("returned correct: %v", string(message))
	}

}

func TestWsWrongEvengtMessage(t *testing.T) {
	testMessage := Event{
		Name: "WrongEvent",
		Data: `this is test data`,
	}
	err := ws.WriteMessage(1, testMessage.Raw())

	if err != nil {
		//t.Fatalf(`Sendmessage got : %v,  "", error`, err)
		t.Fatalf(`Sendmessage got error : %v  `, err)

	} else {
		t.Log("testWS succeded")
	}

	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("error reading reply message: %v", err)
	} else {
		expextedReply := `{"name":"error","error":"Websocket reader received unknown event or message from client"}`
		t.Logf("Msg:%v:", string(message))
		t.Logf("Exp:%v:", expextedReply)
		if string(message) != string(expextedReply) {
			t.Fatalf("expected response : %v , got isntead: %v", expextedReply, string(message))
		}

		t.Logf("returned correct: %v", string(message))
	}

}
func TestWsEchoMessageNested(t *testing.T) {

	type tstData struct {
		Name string `json:Name`
		Data struct {
			Game   int16  `json:Value`
			Status string `json:Value`
		}
	}

	tst := tstData{}
	tst.Name = "TestMessageReply"
	tst.Data.Game = 1
	tst.Data.Status = "Init"

	raw, _ := json.Marshal(tst)

	err := ws.WriteMessage(1, raw)

	if err != nil {
		//t.Fatalf(`Sendmessage got : %v,  "", error`, err)
		t.Fatalf(`Sendmessage got error : %v  `, err)

	} else {
		t.Log("testWS succeded")
	}

	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("error reading reply message: %v", err)
	} else {
		expextedReply := `{"event":"response","data":"THIS IS TEST DATA"}`
		if string(message) != expextedReply {
			t.Fatalf("expected response : %v , got isntead: %v", expextedReply, string(message))
		}

		t.Logf("returned correct: %v", string(message))
	}

}
