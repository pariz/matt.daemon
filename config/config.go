package config

import "github.com/BurntSushi/toml"

type Config struct {
	PidDir string `toml:"pidDir"`
	LogDir string `toml:"logDir"`

	Processes map[string]Process `toml:"process"`
}

type Process struct {
	Description string

	Script string
}

func Load(cfgPath string) (config Config, err error) {

	if _, err = toml.DecodeFile(cfgPath, &config); err != nil {
		return
	}

	return
}
