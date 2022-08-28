package main

import (
	"log"
	"time"

	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
)

func main() {
	cfg, err := MakeConfig("config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	var sender utils.Sender[schema.Notification] = utils.NewAmqMessenger[schema.Notification](
		utils.AmqConfig{
			Endpoint: cfg.ActiveMQ.Endpoint,
			User:     cfg.ActiveMQ.Username,
			Pass:     cfg.ActiveMQ.Password,
		},
		cfg.NotificationAddresses,
		utils.Anycast,
	)

	msgFactory := utils.MakeMessageFactory[schema.Notification]("1.0.0")
	for {
		sender.Send(msgFactory.Make(schema.Notification{
			Type:   schema.Hello,
			Tokens: map[string]string{"name": "alex1"},
		}))
		time.Sleep(1 * time.Second)
		sender.Send(msgFactory.Make(schema.Notification{
			Type:   schema.Bye,
			Tokens: map[string]string{"name": "alex"},
		}))
		time.Sleep(1 * time.Second)
	}
}
