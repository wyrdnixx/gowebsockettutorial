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

func TestMain(m *testing.M) {
	fmt.Println("runiing testmain...")
	go main()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

var addrCLI = flag.String("addrCLI", "localhost:8080", "http service address")

func TestServer(t *testing.T) {

	fmt.Println("runiing...")

	testMessage := Event{
		Name: "message2",
		Data: `{
			"build":"0.1",
			"testdata": "47"
			}`,
	}

	u := url.URL{Scheme: "ws", Host: *addrCLI, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	err = ws.WriteMessage(1, testMessage.Raw())
	if err != nil {
		//t.Fatalf(`Sendmessage got : %v,  "", error`, err)
		t.Fatalf(`Sendmessage got error : %v  `, err)

	}

}
