package types

import (
	"time"

	"github.com/google/uuid"
)

type Metadata map[string]interface{}

type Message struct {
	ID        string      `json:"id"`
	Provider  string      `json:"provider"`
	Timestamp time.Time   `json:"timestamp"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload"`
	RawLog    string      `json:"raw_log"`
	Processed string      `json:"-"`
}

func BuildMessage(provider, raw string) Message {
	return Message{
		ID:        uuid.New().String(),
		Provider:  provider,
		Timestamp: time.Now(),
		Level:     "info",
		Payload:   nil,
		RawLog:    raw,
	}
}
