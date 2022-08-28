package main

import (
	"github.com/alex529/activemq/schema"
	"github.com/alex529/activemq/utils"
)

type Config struct {
	TemplateLinks struct {
		Emails   []schema.NotificationType `yaml:"emails"`
		Webhooks []schema.NotificationType `yaml:"webhooks"`
	} `yaml:"template_links"`
	Webhooks map[string][]string `yaml:"webhooks"`
}

func MakeConfig(filePath string) (Config, error) {
	return utils.MakeConfig[Config](filePath)
}
