package utils

import (
	"encoding/json"
	"log"

	"github.com/alex529/activemq/schema"
	"github.com/go-stomp/stomp"
)

type AmqCastType string

const (
	Anycast   AmqCastType = "ANYCAST"
	Multicast AmqCastType = "MULTICAST"
)

type AmqConfig struct {
	Endpoint, User, Pass string
	Consumer             schema.ConsumerType
}

type AmqMessenger[T any] struct {
	cfg   AmqConfig
	addr  string
	cType AmqCastType
	con   *stomp.Conn
}

func NewAmqMessenger[T any](cfg AmqConfig, address string, castType AmqCastType) *AmqMessenger[T] {
	return &AmqMessenger[T]{
		cfg:   cfg,
		addr:  address,
		cType: castType,
	}
}

func (m *AmqMessenger[T]) connect() error {
	if m.con != nil {
		return nil
	}

	con, err := stomp.Dial(
		"tcp",
		m.cfg.Endpoint,
		stomp.ConnOpt.Login(
			m.cfg.User,
			m.cfg.Pass,
		),
	)
	m.con = con

	return err
}

func (m *AmqMessenger[T]) Disconnect() error {
	if m.con == nil {
		return nil
	}
	return m.con.Disconnect()
}

func (m *AmqMessenger[T]) Send(msg schema.Message[T]) error {
	if err := m.connect(); err != nil {
		return err
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return m.con.Send(m.addr, string(m.cType), []byte(payload))
}

func (m *AmqMessenger[T]) Process(processor func(schema.Message[T], error) error) error {
	if err := m.connect(); err != nil {
		return err
	}

	sub, err := m.con.Subscribe(m.addr, stomp.AckAuto)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		payload := <-sub.C
		var msg schema.Message[T]
		if err := json.Unmarshal(payload.Body, &msg); err != nil {
			log.Printf("Error unmarshalling err: %v\nmsg: %s", err, payload.Body)
			continue
		}

		if !ShouldProcessMessage(m.cfg.Consumer, msg) {
			continue
		}

		if err := processor(msg, payload.Err); err != nil {
			return err
		}
	}
}
