package provider

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

var httpDockerClient = &http.Client{
	Transport: &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", "/var/run/docker.sock")
		},
	},
}

type DockerProvider struct {
	name        string
	containerID string
}

func NewDockerProvider(name, containerID string) *DockerProvider {
	return &DockerProvider{
		name:        name,
		containerID: containerID,
	}
}

func (p *DockerProvider) Stream(ctx context.Context, send chan<- types.Message, metadata types.Metadata) error {
	var since int64
	var status bool

	if s, ok := metadata["since"]; ok {
		t, ok := s.(time.Time)
		if !ok {
			return errors.New("invalid since")
		}
		since = t.Unix()
	}

	url := fmt.Sprintf("http://unix/containers/%s/logs?stdout=true&follow=true&since=%d", p.containerID, since)
	response, err := httpDockerClient.Get(url)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(response.Body)

loop:
	for {
		select {
		case <-ctx.Done():
			log.Println("closing stream")
			_ = response.Body.Close()
			break loop
		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if errors.Is(err, io.EOF) && status {
					metadata["since"] = time.Now()
					send <- types.BuildMessage(p.name, fmt.Sprintf("connection with container %s was closed.", p.containerID))
				}
				return err
			}

			if _, ok := metadata["since"]; ok && !status {
				delete(metadata, "since")
			}

			status = true

			send <- types.BuildMessage(p.name, string(line))
		}
	}

	return nil
}
