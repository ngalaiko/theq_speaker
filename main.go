package main

import (
	"./speaker"
	"./theq_speak"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	configFileName = "config.yaml"
)

func main() {
	config := theq_speak.Config{}
	if err := readConfig(&config); err != nil {
		panic(err)
	}

	speaker := speaker.New(config.ApiKey)
	reader := theq_speak.New(speaker)
	reader.Start()
}

func readConfig(config *theq_speak.Config) error {
	bytes, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, config)
}
