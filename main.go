package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	hostname             string
	port                 string
	socketLocationRemote string
	socketLocationLocal  string
	sshUserName          string
	sshControlSocket     string
	command              string
	configFile           string
)

func init() {
	log.SetFlags(0)
	load_config_file()
	parseflags()
	if socketLocationLocal == "/tmp/.docker-hostname.sock" {
		socketLocationLocal = fmt.Sprintf("/tmp/.docker-%s@%s:%s.sock", sshUserName, hostname, port)
	}
}

func parseflags() {
	// Parse program parameters
	flag.StringVar(&hostname, "host", "localhost", "Hostname that the docker daemon runs on (Default makes no sense :) )")
	flag.StringVar(&port, "port", "22", "SSH port on the remote docker host")
	flag.StringVar(&socketLocationRemote, "socketLocationRemote", "/var/run/docker.sock", "Docker socket location on the remote host")
	flag.StringVar(&socketLocationLocal, "socketLocationsLocal", "/tmp/.docker-hostname.sock", "Docker socket location")
	flag.StringVar(&sshUserName, "ssh-user", os.Getenv("USER"), "SSH username to use on remote server - Leave empty to use your current login")
	flag.StringVar(&sshControlSocket, "ssh-control-socket", "", "Path to the generated SSH control socket - Leave empty to use OpenSSHs config")
	flag.StringVar(&configFile, "config-file", "./.docker-socket-tunnel-helper", "Location of config file")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	log.Printf("Usage: %s need a command and [possibly] options; new, connect or disconnect\n",
		path.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		// fmt.Fprintf(os.Stderr, "Need a command and [possibly] options; new, connect or disconnect\n")
		flag.Usage()
	}

	switch {
	case command == "new":
		setup_config()
	case command == "connect":
		connect_to_server()
	case command == "disconnect":
		disconnect_from_server()
	}

	println(fmt.Sprintf("Making socket %s from host %s:%s available as %s", socketLocationRemote, hostname, port, socketLocationLocal))

}
