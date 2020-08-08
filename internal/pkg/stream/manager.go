package stream

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type resourceProvider interface {
	Stream(context.Context, chan<- types.Message, types.Metadata) error
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
		go m.startProviderStream(ctx, r, sendChan, map[string]interface{}{})
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

func (m *Manager) startProviderStream(ctx context.Context, r resourceProvider, sendChan chan types.Message, meta types.Metadata) {
	if err := r.Stream(ctx, sendChan, meta); err != nil {
		if err == io.EOF {
			time.Sleep(5 * time.Second)
			go m.startProviderStream(ctx, r, sendChan, meta)
			return
		}

		log.Printf("Failed to stream : %s \n", err.Error())
		return
	}
}
