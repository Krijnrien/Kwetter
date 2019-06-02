package main

import (
	"time"
)

const (
	KindKweetCreated = iota + 1
)

type KweetCreatedMessage struct {
	Kind      uint32    `json:"kind"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func newKweetCreatedMessage(id string, body string, createdAt time.Time) *KweetCreatedMessage {
	return &KweetCreatedMessage{
		Kind:      KindKweetCreated,
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
