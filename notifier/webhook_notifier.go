package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
)

type WebhookNotifier struct {
	cfg       Config
	tProvider TemplateProvider
	uProvider UriProvider
}

func MakeWebhookNotifier(cfg Config, tProvider TemplateProvider) WebhookNotifier {
	return WebhookNotifier{
		cfg:       cfg,
		tProvider: tProvider,
		uProvider: MakeUriProvider(cfg),
	}
}

func (n WebhookNotifier) Notify(notification schema.Notification) error {
	payload, err := n.tProvider.Get(Webhook, notification)
	if err != nil {
		return err
	}

	aggErr := utils.NewAggregateError()
	for _, uri := range n.uProvider.GetUris(notification.Type) {
		//todo make async
		aggErr.Add(sendNotification(uri, payload))
	}

	return aggErr.GetError()
}

func sendNotification(uri, payload string) error {
	resp, err := http.Post(
		uri,
		"application/json",
		bytes.NewBuffer([]byte(payload)),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf(string(body))
	return nil
}
