package main

import (
	"io/ioutil"
	"log"

	"github.com/autoai-incubator/notihub/components"
	"gopkg.in/yaml.v2"
)

func main() {
	var config Config
	dat, err := ioutil.ReadFile("./config.yaml")
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if config.Telegram {
		log.Println("Telegram: Enabled")
		b := components.NewTelegramBot()
		go b.Start()
	}
	router := NewRouter()
	router.Run()
}
