package schema

type Message struct {
	Version string `json:"version"`
	Payload []byte `json:"payload"`
}
