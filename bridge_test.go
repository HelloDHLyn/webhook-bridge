package bridge

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	err := LoadConfiguration("./examples/config.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if config.Version != "1" ||
		config.InputPathPrefix != "/webhook" ||
		config.InputPort != 8080 ||
		config.Bridges[0].Input.Source != "docker-hub" ||
		config.Bridges[0].Output.Target != "slack" ||
		config.Bridges[0].Output.Options["channel"] != "#general" {
		t.Fail()
	}
}
