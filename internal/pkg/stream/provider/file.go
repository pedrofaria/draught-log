package provider

import (
	"context"

	"github.com/hpcloud/tail"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type FileProvider struct {
	name string
	path string
}

func NewFileProvider(name, path string) *FileProvider {
	return &FileProvider{
		name: name,
		path: path,
	}
}

func (p *FileProvider) Stream(ctx context.Context, send chan<- types.Message, _ types.Metadata) error {
	t, err := tail.TailFile(p.path, tail.Config{Follow: true})
	if err != nil {
		return err
	}

loop:
	for {
		select {
		case <-ctx.Done():
			t.Done()
			break loop
		case line := <-t.Lines:
			if line.Err != nil {
				t.Done()
				return err
			}
			send <- types.BuildMessage(p.name, line.Text)
		}
	}

	return nil
}
