package main

import (
	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
)

type Config struct {
	ActiveMQ struct {
		Endpoint string `yaml:"endpoint" envconfig:"AMQ_ENDPOINT"`
		Username string `yaml:"username" envconfig:"AMQ_USER"`
		Password string `yaml:"password" envconfig:"AMQ_PASSWORD"`
	} `yaml:"activemq"`
	Subscription  string `yaml:"subscription" envconfig:"SUB_ADR"`
	TemplateLinks struct {
		Emails   []schema.NotificationType `yaml:"emails"`
		Webhooks []schema.NotificationType `yaml:"webhooks"`
	} `yaml:"template_links"`
	Webhooks map[string][]string `yaml:"webhooks"`
}

func MakeConfig(filePath string) (Config, error) {
	return utils.MakeConfig[Config](filePath)
}
