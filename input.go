package bridge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type InputSource interface {
	GetInput(r *http.Request) (*Input, error)
}

func GetInputSource(source string) InputSource {
	if source == "docker-hub" {
		return &DockerHubInputSource{}
	}
	return nil
}

type DockerHubInputSource struct {
	options map[string]string
}

func (src *DockerHubInputSource) GetInput(r *http.Request) (*Input, error) {
	if r.Method != "POST" {
		return nil, fmt.Errorf("invalid http method: " + r.Method)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Input{Source: src, CalledAt: &now, Payload: body}, nil
}

type Input struct {
	Source   InputSource
	CalledAt *time.Time
	Payload  []byte
}
