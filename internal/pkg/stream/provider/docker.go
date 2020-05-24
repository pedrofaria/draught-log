package provider

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

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

func (p *DockerProvider) Stream(ctx context.Context, send chan<- types.Message) error {
	url := fmt.Sprintf("http://unix/containers/%s/logs?stdout=true&follow=true", p.containerID)
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
			break loop
		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				return err
			}

			newLine := strings.Trim(string(line), "\r\n \u0000\u0001")

			send <- types.BuildMessage(p.name, newLine)
		}
	}

	return nil
}
