package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/pariz/matt.daemon/config"
	"github.com/pariz/matt.daemon/process"
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "config", "/etc/mattdaemon.toml", "Provide the config path for matt.daemon")

}

func main() {
	flag.Parse()

	// Show epic ascii art
	fmt.Println(matt)

	// Load config
	cfg, err := config.Load(configFilePath)

	if err != nil {
		panic(err)
	}

	// Spawn processes
	process.Init(cfg)

	// Setup Rpc
	rpcServer := new(process.Rpc)
	rpc.Register(rpcServer)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// Make sure program doesn't terminate
	for {
	}
}
