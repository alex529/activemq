package main

import "github.com/alex529/activemq/schema"

type UriProvider struct {
	cfg Config
}

func MakeUriProvider(cfg Config) UriProvider {
	return UriProvider{cfg: cfg}
}

func (p UriProvider) GetUris(nType schema.NotificationType) []string {
	return p.cfg.Webhooks[string(nType)]
}
