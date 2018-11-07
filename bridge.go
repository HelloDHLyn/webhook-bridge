package bridge

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

var (
	converterMap = make(map[string](*converter))

	config       configuration
	configLoaded = false
)

// Configuration contains all information for bridging.
type configuration struct {
	Version string

	InputPathPrefix string `yaml:"input_path_prefix"`
	InputPort       int    `yaml:"input_port"`

	Bridges []struct {
		Name  string
		Input struct {
			Source  string
			Options map[string]string
		}
		Output struct {
			Target  string
			Options map[string]string
		}
		Converter struct {
			Name string
		}
	}
}

// Converter accept Input as argument, and should returns Output.
type converter func(interface{}) (interface{}, error)

func LoadConfiguration(filePath string) error {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(body, &config)
	if err != nil {
		return err
	}

	configLoaded = true
	return nil
}

// Register a input-output converter.
func RegisterConverter(name string, conv converter) {
	converterMap[name] = &conv
}

// Start HTTP server to accept webhooks.
func Start() error {
	if !configLoaded {
		return fmt.Errorf("configuration not loaded")
	}

	for _, b := range config.Bridges {
		pattern := config.InputPathPrefix + "/" + b.Name
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

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.InputPort), nil))
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
