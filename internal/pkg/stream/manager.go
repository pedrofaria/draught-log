package stream

import (
	"context"
	"log"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/provider"
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
	IsDockerEnabled bool
	FileDirectory   string

	providers []resourceProvider
}

// GetListResources returns a list of all available resources of each provider
func (m *Manager) GetListResources() ([]Resource, error) {
	var list []Resource

	if m.IsDockerEnabled {
		dockerRes, err := provider.GetDockerResources()
		if err != nil {
			return nil, err
		}

		for _, re := range dockerRes {
			if re.State != "running" {
				continue
			}

			list = append(list, Resource{
				ID:       re.Id,
				Name:     re.Names[0],
				Provider: "docker",
			})
		}
	}

	if m.FileDirectory != "" {
		fileRes, err := provider.GetFileResources(m.FileDirectory)
		if err != nil {
			return nil, err
		}

		for _, re := range fileRes {
			list = append(list, Resource{
				ID:       re.Path,
				Name:     re.Path,
				Provider: "file",
			})
		}
	}

	return list, nil
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
