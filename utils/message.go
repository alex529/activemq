package utils

import "github.com/alex529/activemq/schema"

type MessageFactory[T any] struct {
	version string
	msg     T
}

func MakeMessageFactory[T any](version string) MessageFactory[T] {
	return MessageFactory[T]{
		version: version,
	}
}

func (p MessageFactory[T]) Make(payload T) schema.Message[T] {
	return schema.Message[T]{
		Version: p.version,
		Payload: payload,
	}
}

func (p MessageFactory[T]) MakeWithConsumer(consumers []schema.ConsumerType, payload T) schema.Message[T] {
	return schema.Message[T]{
		Consumers: consumers,
		Version:   p.version,
		Payload:   payload,
	}
}
