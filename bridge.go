package bridge

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	converterMap = make(map[string](*converter))
)

// Converter accept Input as argument, and should returns Output.
type converter func(interface{}) (interface{}, error)

// Register a input-output converter.
func RegisterConverter(name string, conv converter) {
	converterMap[name] = &conv
}

// Start HTTP server to accept webhooks.
func Start() error {
	if loadedConfig == nil {
		return fmt.Errorf("configuration not loaded")
	}

	for _, b := range loadedConfig.Bridges {
		pattern := loadedConfig.Server.PathPrefix + "/" + b.Name
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			input, err := parseInput(b.Input.Source, r)

			conv := *converterMap[b.Converter.Name]
			output, err := conv(input)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = output.(Output).Send(b.Output.Options)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
		})
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(loadedConfig.Server.Port), nil))
	return nil
}

func parseInput(source string, r *http.Request) (Input, error) {
	if source == "docker-hub" {
		var input DockerHubInput
		err := input.ParseHTTPRequest(r)
		return &input, err
	}
	return nil, fmt.Errorf("invalid source name")
}
