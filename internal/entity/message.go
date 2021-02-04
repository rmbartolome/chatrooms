package entity

import "context"

type Message struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message Message) error
	GetMessage(ctx context.Context, id string) (string, error)
}
