package schema

import "strings"

type NotificationType string

const (
	Hello NotificationType = "hello"
)

var (
	notificationMap = map[string]NotificationType{
		"hello": Hello,
	}
)

func ParseString(str string) (NotificationType, bool) {
	c, ok := notificationMap[strings.ToLower(str)]
	return c, ok
}

type Notification struct {
	Type   string
	Tokens map[string]string
}
