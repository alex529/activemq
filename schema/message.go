package schema

type Message[T any] struct {
	Version string `json:"version"`
	Payload T      `json:"payload"`
}
