package schema

import "strings"

type NotificationType string

const (
	Hello NotificationType = "hello"
	Bye   NotificationType = "bye"
)

var (
	notificationMap = map[string]NotificationType{
		"hello": Hello,
		"bye":   Bye,
	}
)

func ParseString(str string) (NotificationType, bool) {
	c, ok := notificationMap[strings.ToLower(str)]
	return c, ok
}

type Notification struct {
	Type   NotificationType
	Tokens map[string]string
}
