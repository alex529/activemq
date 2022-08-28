package main

import "github.com/alex529/activemq/utils"

type Config struct {
	ActiveMQ struct {
		Endpoint string `yaml:"endpoint" envconfig:"AMQ_ENDPOINT"`
		Username string `yaml:"username" envconfig:"AMQ_USER"`
		Password string `yaml:"password" envconfig:"AMQ_PASSWORD"`
	} `yaml:"activemq"`
	NotificationAddresses string `yaml:"NotificationAddresses" envconfig:"NOTIFICATIONS_ADR"`
}

func MakeConfig(filePath string) (Config, error) {
	return utils.MakeConfig[Config](filePath)
}
