package utils

import (
	"github.com/alex529/activemq/schema"
	"golang.org/x/exp/slices"
)

type Sender[T any] interface {
	Send(msg schema.Message[T]) error
}

type Processor[T any] interface {
	Process(func(schema.Message[T], error) error) error
}

type Messenger[T any] interface {
	Sender[T]
	Processor[T]
}

func ShouldProcessMessage[T any](consumer schema.ConsumerType, msg schema.Message[T]) bool {
	if msg.Consumers == nil || len(msg.Consumers) == 0 {
		return true
	}
	if slices.Contains(msg.Consumers, schema.None) {
		return false
	}
	return slices.Contains(msg.Consumers, consumer)
}
