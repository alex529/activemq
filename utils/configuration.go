package utils

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func MakeConfig[T any](filePath string) (T, error) {
	var cfg T
	if err := readFile(filePath, &cfg); err != nil {
		return cfg, err
	}

	return cfg, envconfig.Process("", &cfg)
}

func readFile[T any](path string, cfg *T) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(cfg)
}
