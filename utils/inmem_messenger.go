package utils

import (
	"time"

	"github.com/alex529/activemq/schema"
)

type InMemoryMessenger[T any] struct {
	q              []schema.Message[T]
	stopProcessing bool
}

func MakeInMemoryMessenger[T any]() InMemoryMessenger[T] {
	return InMemoryMessenger[T]{
		q: make([]schema.Message[T], 0),
	}
}

func (m InMemoryMessenger[T]) Send(msg schema.Message[T]) error {
	m.q = append(m.q, msg)
	return nil
}

func (m InMemoryMessenger[T]) Process(processor func(msg schema.Message[T]) error) error {
	for !m.stopProcessing {
		if len(m.q) < 1 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		msg := m.q[len(m.q)-1]
		if err := processor(msg); err != nil {
			return err
		}

		m.q = m.q[:len(m.q)-1]
	}
	return nil
}
