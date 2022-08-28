package main

import "github.com/alex529/activemq/utils"

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"PORT"`
	} `yaml:"server"`
}

func MakeConfig(filePath string) (Config, error) {
	return utils.MakeConfig[Config](filePath)
}
