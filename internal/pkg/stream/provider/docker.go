package provider

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type DockerResource struct {
	Id    string
	Names []string
}

func GetDockerResources() ([]DockerResource, error) {
	response, err := httpDockerClient.Get("http://unix/containers/json")
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var l []DockerResource

	if err := json.Unmarshal(buf, &l); err != nil {
		return nil, err
	}

	return l, nil
}

type DockerProvider struct {
	containerID string
}

func NewDockerProvider(containerID string) *DockerProvider {
	return &DockerProvider{
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

			send <- types.BuildMessage(p.containerID, newLine)
		}
	}

	return nil
}
