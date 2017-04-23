package main

import (
	"flag"
	"fmt"
	"github.com/ngalayko/theq_ask/speaker"
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
	s.Start()
}
