package main

import (
	"log"

	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
)

func main() {
	cfg, err := MakeConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	tProvider, err := MakeTemplateProvider("./templates/email", "./templates/webhook") //todo
	if err != nil {
		log.Fatal(err)
	}
	dist := MakeDistributor(cfg, tProvider)

	var p utils.Processor[schema.Notification] = utils.NewAmqMessenger[schema.Notification](
		utils.AmqConfig{
			Endpoint: cfg.ActiveMQ.Endpoint,
			User:     cfg.ActiveMQ.Username,
			Pass:     cfg.ActiveMQ.Password,
		},
		cfg.Subscription,
		utils.Anycast,
	)

	err = p.Process(func(msg schema.Message[schema.Notification], err error) error {
		if err := dist.NotifyAll(msg.Payload); err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
