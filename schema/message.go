package schema

type ConsumerType int

const (
	None ConsumerType = iota
	All
	Notifier
)

type Message[T any] struct {
	Version   string         `json:"version"`
	Consumers []ConsumerType `json:"consumer"`
	Payload   T              `json:"payload"`
}
