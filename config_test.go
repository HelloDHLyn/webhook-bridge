package bridge

import (
	"testing"
)

func TestLoadConfigurationFromFile(t *testing.T) {
	err := LoadConfigurationFromFile("./examples/config.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if loadedConfig.Version != "1" ||
		loadedConfig.Server.PathPrefix != "/webhook" ||
		loadedConfig.Server.Port != 8080 ||
		loadedConfig.Bridges[0].Input.Source != "docker-hub" ||
		loadedConfig.Bridges[0].Output.Target != "slack" {
		t.Fail()
	}
}

func TestLoadConfigurationFromHTTP(t *testing.T) {
	err := LoadConfigurationFromHTTP("https://gist.githubusercontent.com/HelloDHLyn/17c9f3401fe7e06f4a993c66affc77ee/raw/8a1d793c294ee1806e74bdc0198480c09c563866/config.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if loadedConfig.Version != "1" ||
		loadedConfig.Server.PathPrefix != "/webhook" ||
		loadedConfig.Server.Port != 8080 ||
		loadedConfig.Bridges[0].Input.Source != "docker-hub" ||
		loadedConfig.Bridges[0].Output.Target != "slack" {
		t.Fail()
	}
}
