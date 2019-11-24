package bridge

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

type OutputTarget interface {
	SendOutput(output *Output) error
}

func GetOutputTarget(target string, options map[string]string) OutputTarget {
	if target == "slack" {
		return &SlackOutputTarget{options}
	}
	return nil
}

type SlackOutputTarget struct {
	options map[string]string
}

func (tgt *SlackOutputTarget) SendOutput(o *Output) error {
	url := tgt.options["url"]
	if url == "" {
		return fmt.Errorf("slack output requires `url` option")
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(o.ConvertedPayload))
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

type Output struct {
	ConvertedPayload []byte
}
