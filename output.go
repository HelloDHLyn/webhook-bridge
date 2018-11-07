package bridge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Output interface {
	Send(options map[string]string) error
}

/// Slack outgoing webhook.
/// Reference: https://api.slack.com/incoming-webhooks
type SlackOutput struct {
	Text        string                  `json:"text"`
	Attachments []SlackOutputAttachment `json:"attachments"`
}

type SlackOutputAttachment struct {
	Title          string                        `json:"title"`
	Fields         []SlackOutputAttachmentField  `json:"fields,omitempty"`
	AuthorName     string                        `json:"author_name,omitempty"`
	AuthorIcon     string                        `json:"author_icon,omitempty"`
	ImageURL       string                        `json:"image_url,omitempty"`
	Text           string                        `json:"text,omitempty"`
	Fallback       string                        `json:"fallback,omitempty"`
	CallbackID     string                        `json:"callback_id,omitempty"`
	Color          string                        `json:"color,omitempty"`
	AttachmentType string                        `json:"attachment_type,omitempty"`
	Actions        []SlackOutputAttachmentAction `json:"actions,omitempty"`
}

type SlackOutputAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type SlackOutputAttachmentAction struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Send slack webhook.
//
// Options:
//   - `url`(required) : Webhook URL.
func (p *SlackOutput) Send(options map[string]string) error {
	url := options["url"]
	if url == "" {
		return fmt.Errorf("slack output requires `url` option")
	}

	body, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("server returned status " + strconv.Itoa(res.StatusCode))
	}

	return nil
}
