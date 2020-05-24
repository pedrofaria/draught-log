package provider

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/hpcloud/tail"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type FileResource struct {
	Path string
}

func GetFileResources(dir string) ([]FileResource, error) {
	var r []FileResource

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if path == dir {
				return nil
			}

			r = append(r, FileResource{Path: strings.TrimPrefix(path, dir)})
			return nil
		})
	if err != nil {
		return nil, err
	}

	return r, nil
}

type FileProvider struct {
	dir  string
	path string
}

func NewFileProvider(dir, path string) *FileProvider {
	return &FileProvider{
		dir:  dir,
		path: path,
	}
}

func (p *FileProvider) Stream(ctx context.Context, send chan<- types.Message) error {
	path := p.dir + p.path

	t, err := tail.TailFile(path, tail.Config{Follow: true})
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
			send <- types.BuildMessage(p.path, line.Text)
		}
	}

	return nil
}
