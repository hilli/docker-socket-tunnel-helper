package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	SshHostname          string
	SshPort              string
	SocketLocationRemote string
	SocketLocationLocal  string
	SshUserName          string
	SshControlSocket     string
	Command              string
	ConfigFile           string
}

var (
	conf Config
)

func init() {
	log.SetFlags(0)
	setflags()
	load_config_file()
	flag.Parse() // Load again to override config file with cmdline params
	if conf.SocketLocationLocal == "/tmp/.docker-hostname.sock" {
		conf.SocketLocationLocal = fmt.Sprintf("/tmp/.docker-%s@%s:%s.sock", conf.SshUserName, conf.SshHostname, conf.SshPort)
	}
}

func setflags() {
	// Parse program parameters
	flag.StringVar(&conf.SshHostname, "host", "localhost", "Hostname that the docker daemon runs on (Default makes no sense :) )")
	flag.StringVar(&conf.SshPort, "port", "22", "SSH port on the remote docker host")
	flag.StringVar(&conf.SocketLocationRemote, "socketLocationRemote", "/var/run/docker.sock", "Docker socket location on the remote host")
	flag.StringVar(&conf.SocketLocationLocal, "socketLocationsLocal", "/tmp/.docker-hostname.sock", "Docker socket location")
	flag.StringVar(&conf.SshUserName, "ssh-user", os.Getenv("USER"), "SSH username to use on remote server - Leave empty to use your current login")
	flag.StringVar(&conf.SshControlSocket, "ssh-control-socket", "", "Path to the generated SSH control socket - Leave empty to use OpenSSHs config")
	flag.StringVar(&conf.ConfigFile, "config-file", "./.docker-socket-tunnel-helper", "Location of config file")
	flag.Usage = usage
}

func usage() {
	log.Printf("Usage: %s need a command and [possibly] options; new, connect or disconnect\n",
		path.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	if len(os.Args) > 1 {
		conf.Command = os.Args[1]
	} else {
		// fmt.Fprintf(os.Stderr, "Need a command and [possibly] options; new, connect or disconnect\n")
		flag.Usage()
	}

	switch {
	case conf.Command == "new":
		setup_config()
	case conf.Command == "connect":
		connect_to_server()
	case conf.Command == "disconnect":
		disconnect_from_server()
	}

	println(fmt.Sprintf("Making socket %s from host %s:%s available as %s", conf.SocketLocationRemote, conf.SshHostname, conf.SshPort, conf.SocketLocationLocal))
	println(fmt.Sprintf("Config file: %s", conf.ConfigFile))

}
