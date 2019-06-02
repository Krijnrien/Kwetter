package event

import (
	"time"
)

type Message interface {
	Key() string
}

type KweetCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (m *KweetCreatedMessage) Key() string {
	return "Kweet.created"
}
