package main

import (
	"bridge"
)

func main() {
	bridge, err := bridge.NewService("./examples/config.yaml")
	if err != nil {
		panic(err)
	}

	bridge.StartServer()
}
