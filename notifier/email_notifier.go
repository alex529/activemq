package main

import (
	"log"

	"github.com/alex529/activemq/schema"
)

type EmailNotifier struct {
	cfg          Config
	tProvider    TemplateProvider
	subsProvider EmailSubscriberProvider
}

func MakeEmailNotifier(cfg Config, tProvider TemplateProvider) EmailNotifier {
	return EmailNotifier{
		cfg:          cfg,
		tProvider:    tProvider,
		subsProvider: MakeEmailSubscriberProvider(),
	}
}

func (n EmailNotifier) Notify(notification schema.Notification) error {
	payload, err := n.tProvider.Get(Email, notification)
	if err != nil {
		return err
	}

	for _, to := range n.subsProvider.GetSubs(notification.Type) {
		sendEmail(to, string(notification.Type), payload)
	}

	return nil
}

func sendEmail(to, subject, msg string) {
	log.Printf("Email sender to: %s subject: %s msg: \n%s\n", to, subject, msg)
}
