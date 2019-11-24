package bridge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		PathPrefix string
		Port       int
	}

	Bridges []Bridge
}

// configYAML contains all information for bridging.
type configYAML struct {
	Version string

	Server struct {
		PathPrefix string `yaml:"path_prefix"`
		Port       int    `yaml:"port"`
	}

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
			JSON string
		}
	}
}

// newConfigFromFile reads yaml file and load it.
func newConfigFromFile(filePath string) (*Config, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return parseYAML(body)
}

// newConfigFromHTTP downloads yaml file with HTTP GET request from the URL,
// and load it as configuration.
func newConfigFromHTTP(url string) (config *Config, err error) {
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return
	}
	if res.StatusCode >= 300 {
		err = fmt.Errorf("server returned status " + strconv.Itoa(res.StatusCode))
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return parseYAML(body)
}

func parseYAML(body []byte) (config *Config, err error) {
	var cfgYAML configYAML
	err = yaml.Unmarshal(body, &cfgYAML)
	if err != nil {
		return
	}

	var bridges []Bridge
	for _, b := range cfgYAML.Bridges {
		input := GetInputSource(b.Input.Source)
		output := GetOutputTarget(b.Output.Target, b.Output.Options)

		var converter Converter
		if b.Converter.JSON != "" {
			converter = NewJSONConverter([]byte(b.Converter.JSON))
		}

		bridges = append(bridges, NewBridge(b.Name, input, output, converter))
	}

	return &Config{
		Server: struct {
			PathPrefix string
			Port       int
		}{cfgYAML.Server.PathPrefix, cfgYAML.Server.Port},
		Bridges: bridges,
	}, nil
}
