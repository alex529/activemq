package utils

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func MakeConfig[T any](filePath string) (T, error) {
	var cfg T
	if err := readFile(filePath, &cfg); err != nil {
		return cfg, err
	}
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, err
	}

	return cfg, areAllInitialized(cfg)
}

func readFile[T any](path string, cfg *T) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(cfg)
}

func areAllInitialized(cfg interface{}) error {
	reflectType := reflect.TypeOf(cfg)
	reflectValue := reflect.ValueOf(cfg)

	aggErr := NewAggregateError()
	if reflectType.Kind() == reflect.Array || reflectType.Kind() == reflect.Slice {
		typeName := reflectType.Name()
		valueValue := reflectValue.Interface()
		if reflect.TypeOf(valueValue).Elem().Kind() != reflect.String {
			return errors.New("Only arrays of strings are allowed")
		}
		val := reflect.ValueOf(valueValue)
		if val.Len() == 0 || val.Index(0).Interface() == "" {
			return fmt.Errorf("%s was not set", typeName)
		}
	}
	for i := 0; reflectType.Kind() == reflect.Struct && i < reflectType.NumField(); i++ {
		typeName := reflectType.Field(i).Name
		valueValue := reflectValue.Field(i).Interface()

		tmp := reflectValue.Field(i).Kind()
		switch tmp {
		case reflect.String:
			if valueValue == "" {
				aggErr.Add(fmt.Errorf("%s was not set", typeName))
			}
		case reflect.Struct:
			aggErr.Add(areAllInitialized(valueValue))
		case reflect.Map:
			aggErr.Add(areAllInitialized(valueValue))
		case reflect.Slice:
			aggErr.Add(areAllInitialized(valueValue))
		default:
			return errors.New("Only structs/arrays of strings are allowed")
		}

	}
	return aggErr.GetError()
}
