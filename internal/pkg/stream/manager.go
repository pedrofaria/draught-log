package stream

import (
	"context"
	"log"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type Resource struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Provider string `json:"type"`
}

type resourceProvider interface {
	Stream(context.Context, chan<- types.Message) error
}

type Manager struct {
	providers []resourceProvider
}

func (m *Manager) RegisterProvider(p resourceProvider) {
	m.providers = append(m.providers, p)
}

func (m *Manager) Stream(ctx context.Context, h func(types.Message)) {
	sendChan := make(chan types.Message)

	for _, r := range m.providers {
		go func(r resourceProvider) {
			if err := r.Stream(ctx, sendChan); err != nil {
				log.Printf("Failed to stream : %s \n", err.Error())
				return
			}
		}(r)
	}

loop:
	for {
		select {
		case msg := <-sendChan:
			h(msg)
		case <-ctx.Done():
			break loop
		}
	}
}
