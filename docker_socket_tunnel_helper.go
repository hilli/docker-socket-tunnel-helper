package main

import (
	"fmt"
	"os"
)

func load_config_file() {
	_, err := os.Stat(configFile)
	if err != nil {
		// Config found - Overwriting

	}
}

func setup_config() {

}

func connect_to_server() {
	println("Connecting...")
	println(fmt.Sprintf("export DOCKER_HOST=%s", socketLocationLocal))
}

func disconnect_from_server() {
	println("Disconnecting...")
}
