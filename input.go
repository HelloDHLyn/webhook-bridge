package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input interface {
	ParseHTTPRequest(r *http.Request) error
}

type DockerHubInput struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Images   []string `json:"images"`
		PushedAt float64  `json:"pushed_at"`
		Pusher   string   `json:"pusher"`
		Tag      string   `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int     `json:"comment_count"`
		DateCreated     float64 `json:"date_created"`
		Description     string  `json:"description"`
		Dockerfile      string  `json:"dockerfile"`
		FullDescription string  `json:"full_description"`
		IsOfficial      bool    `json:"is_official"`
		IsPrivate       bool    `json:"is_private"`
		IsTrusted       bool    `json:"is_trusted"`
		Name            string  `json:"name"`
		Namespace       string  `json:"namespace"`
		Owner           string  `json:"owner"`
		RepoName        string  `json:"repo_name"`
		RepoURL         string  `json:"repo_url"`
		StarCount       int     `json:"star_count"`
		Status          string  `json:"status"`
	} `json:"repository"`
}

func (p *DockerHubInput) ParseHTTPRequest(r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("invalid http method: " + r.Method)
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return err
	}

	return nil
}
