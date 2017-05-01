package main

import (
	"github.com/ngalayko/theq_speaker/server/speaker"
	"log"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const (
	configPath = "./config.yaml"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal("ReadConfig error", err)
	}

	s := speaker.New(config)
	go s.Start()

	http.HandleFunc("/", s.ServeWs)

	if err := http.ListenAndServe(config.Listen, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func ReadConfig() (speaker.Config, error) {
	config := speaker.Config{}

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	return config, err
}
