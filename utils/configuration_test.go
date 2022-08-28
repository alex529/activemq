package utils_test

import (
	"testing"

	"github.com/alex529/activemq/utils"
	"github.com/stretchr/testify/assert"
)

type config struct {
	ActiveMQ struct {
		Endpoint string `yaml:"endpoint" envconfig:"AMQ_ENDPOINT"`
		Username string `yaml:"username" envconfig:"AMQ_USER"`
		Password string `yaml:"password" envconfig:"AMQ_PASSWORD"`
	} `yaml:"activemq"`
	Subscription string   `yaml:"subscription" envconfig:"SUB_ADR"`
	Array        []string `yaml:"array" envconfig:"Arr"`
}

func TestMakeConfigFileRead(t *testing.T) {
	cfg, err := utils.MakeConfig[config]("./testData/cfg.yaml")

	assert.Nil(t, err)
	assert.Equal(t, cfg.Subscription, "notifications")
	assert.Equal(t, cfg.ActiveMQ.Endpoint, "172.21.0.2:61616")
	assert.Equal(t, cfg.ActiveMQ.Username, "quarkus")
	assert.Equal(t, cfg.ActiveMQ.Password, "pass")
	assert.Equal(t, len(cfg.Array), 2)
	assert.Equal(t, cfg.Array[0], "notifications")
}

func TestMakeConfig_ThrowsError_OnEmptyString(t *testing.T) {
	_, err := utils.MakeConfig[config]("./testData/cfg_bad_string.yaml")

	assert.NotNil(t, err)
}

func TestMakeConfig_ThrowsError_OnEmptyArray(t *testing.T) {
	_, err := utils.MakeConfig[config]("./testData/cfg_bad_array.yaml")

	assert.NotNil(t, err)
}
