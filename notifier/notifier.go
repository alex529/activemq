package main

import "github.com/alex529/activemq/schema"

type Notifier interface {
	Notify(schema.Notification) error
}
