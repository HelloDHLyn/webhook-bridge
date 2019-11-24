package main

import (
	"bridge"
	"errors"
	"os"
)

func main() {
	configPath := os.Getenv("BRIDGE_CONFIG_PATH")
	if configPath == "" {
		panic(errors.New("BRIDGE_CONFIG_PATH"))
	}

	bridge, err := bridge.NewService(configPath)
	if err != nil {
		panic(err)
	}

	bridge.StartServer()
}
