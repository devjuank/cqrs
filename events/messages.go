package events

import "time"

type Message interface {
	Type() string
}

type CreateFeedMessage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m CreateFeedMessage) Type() string {
	return "created_feed"
}
