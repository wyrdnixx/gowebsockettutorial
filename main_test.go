package main

import (
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

func TestWs(t *testing.T) {
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

// func TestServerReply(t *testing.T) {

// 	fmt.Println("running...")

// 	testMessage := Event{
// 		Name: "TestMessageReply2",
// 		Data: `this is test data`,
// 	}

// 	u := url.URL{Scheme: "ws", Host: *addrCLI, Path: "/ws"}
// 	log.Printf("connecting to %s", u.String())

// 	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = ws.WriteMessage(1, testMessage.Raw())
// 	if err != nil {
// 		//t.Fatalf(`Sendmessage got : %v,  "", error`, err)
// 		t.Fatalf(`Sendmessage got error : %v  `, err)

// 	}
// 	_, message, err := ws.ReadMessage()
// 	if err != nil {
// 		t.Fatalf("error reading reply message: %v", err)
// 	} else {
// 		expextedReply := `{"event":"response","data":"THIS IS TEST DATA"}`
// 		if string(message) != expextedReply {
// 			//t.Fatalf("expected response : %v , got isntead: %v", expextedReply, string(message))
// 		}

// 		log.Printf("returned correct: %v", string(message))
// 	}

// 	//err := json.Unmarshal(message,m)

// }
