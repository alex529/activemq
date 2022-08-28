package main

import (
	"io/ioutil"
	"strings"

	"github.com/aymerick/raymond"

	"github.com/alex529/activemq/schema"
)

type TemplateType int

const (
	Webhook TemplateType = iota
	Email
)

type TemplateProvider struct {
	templates map[TemplateType]map[schema.NotificationType]string
}

func MakeTemplateProvider(emailTemplateFolderPath, webhookTemplateFolderPath string) (TemplateProvider, error) {
	emailMap, err := getTemplates(emailTemplateFolderPath)
	if err != nil {
		return TemplateProvider{}, err
	}
	webhookMap, err := getTemplates(webhookTemplateFolderPath)
	if err != nil {
		return TemplateProvider{}, err
	}
	templates := make(map[TemplateType]map[schema.NotificationType]string)
	templates[Email] = emailMap
	templates[Webhook] = webhookMap

	return TemplateProvider{
		templates: templates,
	}, nil
}

func getTemplates(folderPath string) (map[schema.NotificationType]string, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	templates := make(map[schema.NotificationType]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		nType, ok := schema.ParseString(getFileNameWithoutExtension(file.Name()))
		if !ok {
			continue
		}

		data, err := ioutil.ReadFile(folderPath + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		templates[nType] = string(data)
	}
	return templates, nil
}

func getFileNameWithoutExtension(name string) string {
	return name[:strings.LastIndex(name, ".")]
}

func (p TemplateProvider) Get(tType TemplateType, notification schema.Notification) (string, error) {
	return raymond.Render(p.templates[tType][notification.Type], notification.Tokens)
}
