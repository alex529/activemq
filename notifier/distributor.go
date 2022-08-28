package main

import (
	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
	"golang.org/x/exp/slices"
)

type Distributor struct {
	cfg   Config
	email EmailNotifier
	web   WebhookNotifier
}

func MakeDistributor(cfg Config, tProvider TemplateProvider) Distributor {
	return Distributor{
		cfg:   cfg,
		email: MakeEmailNotifier(cfg, tProvider),
		web:   MakeWebhookNotifier(cfg, tProvider),
	}
}

func (d Distributor) NotifyAll(notification schema.Notification) error {
	aggErr := utils.NewAggregateError()
	if slices.Contains(d.cfg.TemplateLinks.Emails, notification.Type) {
		aggErr.Add(d.email.Notify(notification))
	}
	if slices.Contains(d.cfg.TemplateLinks.Webhooks, notification.Type) {
		aggErr.Add(d.web.Notify(notification))
	}
	return aggErr.GetError()
}
