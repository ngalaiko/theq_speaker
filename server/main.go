package main

import (
	"flag"
	"fmt"
	"github.com/ngalayko/theq_speaker/server/speaker"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

var (
	key = flag.String("key", "", "Yandex SpeechKit API key")
)

func main() {
	flag.Parse()
	if len(*key) == 0 {
		fmt.Println("You should use -key flag")
		return
	}

	s := speaker.New(*key)

	http.Handle("/", websocket.Handler(s.Start))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
