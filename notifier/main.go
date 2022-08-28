package main

import (
	"log"
	"time"

	"github.com/alex529/activemq/schema"
)

func main() {
	cfg, err := MakeConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//todo subscribe

	tProvider, err := MakeTemplateProvider("./templates/email", "./templates/webhook") //todo
	if err != nil {
		log.Fatal(err)
	}
	dist := MakeDistributor(cfg, tProvider)

	for true {
		dist.NotifyAll(schema.Notification{
			Type:   schema.Hello,
			Tokens: map[string]string{"name": "alex"},
		})
		time.Sleep(300 * time.Millisecond)
		dist.NotifyAll(schema.Notification{
			Type:   schema.Bye,
			Tokens: map[string]string{"name": "alex"},
		})
		time.Sleep(300 * time.Millisecond)
	}
}
