package main

import "github.com/alex529/activemq/schema"

type EmailSubscriberProvider struct {
	subs map[schema.NotificationType][]string
}

func MakeEmailSubscriberProvider() EmailSubscriberProvider {
	return EmailSubscriberProvider{
		subs: map[schema.NotificationType][]string{
			schema.Hello: {"alex@b.com"},
		},
	}
}

func (p EmailSubscriberProvider) GetSubs(nType schema.NotificationType) []string {
	return p.subs[nType]
}
