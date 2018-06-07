package main

import (
	"fmt"
)

func load_config_file() {
	// _, err := os.Stat(conf.ConfigFile)
	// if err != nil {
	// 	if _, err := toml.DecodeFile(conf.ConfigFile, &conf); err != nil {
	// 		println("Failed to load config file") // stderr?
	// 	}

	// }
}

func setup_config() {

}

func connect_to_server() {
	println("Connecting...")
	println(fmt.Sprintf("export DOCKER_HOST=%s", conf.SocketLocationLocal))
}

func disconnect_from_server() {
	println("Disconnecting...")
}
