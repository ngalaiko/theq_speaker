package main

import (
	"flag"
	"fmt"
	"github.com/ngalayko/theq_ask/speaker"
	"github.com/ngalayko/theq_ask/theq_speak"
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

	speaker := speaker.New(*key)
	reader := theq_speak.New(speaker)
	reader.Start()
}
