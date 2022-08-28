package utils

import "github.com/alex529/activemq/schema"

type Sender[T any] interface {
	Send(msg schema.Message[T]) error
}

type Processor[T any] interface {
	Process(func(msg schema.Message[T]) error) error
}
