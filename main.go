package main

import (
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
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
	spew.Dump(cfg)
	// Spawn processes
	process.Init(cfg)

	// Make sure program doesn't terminate
	for {
	}
}
