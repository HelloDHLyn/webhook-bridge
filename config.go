package bridge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

var loadedConfig *configuration

// Configuration contains all information for bridging.
type configuration struct {
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
			Name string
		}
	}
}

// LoadConfigurationFromFile reads yaml file and load it.
//
// If you set the value of `BRIDGE_LOAD_CONFIG_FROM_FILE` environment variable
// as a path of the configuration file, this function will called automatically.
func LoadConfigurationFromFile(filePath string) error {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(body, &loadedConfig)
	if err != nil {
		return err
	}

	fmt.Printf("Loaded %d bridge(s) from the file!\n", len(loadedConfig.Bridges))
	return nil
}

// LoadConfigurationFromHTTP downloads yaml file with HTTP GET request from
// the URL, and load it as configuration.
//
// If you set the value of `BRIDGE_LOAD_CONFIG_FROM_HTTP` environment variable
// as a URL to load configurations, this function will called automatically.
func LoadConfigurationFromHTTP(url string) error {
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode >= 300 {
		return fmt.Errorf("server returned status " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(body, &loadedConfig)
	if err != nil {
		return err
	}

	fmt.Println(loadedConfig)

	fmt.Printf("Loaded %d bridge(s) from the http request!\n", len(loadedConfig.Bridges))
	return nil
}
