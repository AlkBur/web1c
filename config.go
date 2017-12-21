package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const fileCFG = "config.yml"

var (
	// Debug var to switch mode from outside
	debug bool
)

type Config struct {
	Debug  bool   `json:"debug"`
	Addr   string `json:"addr"`
	Bases []struct{
		Path string `json:"path"`
		Name string `json:"name"`
		Url string  `json:"url"`
	} `json:"bases"`
}

func parseConfig(file string) (*Config, error) {
	cfg, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := &Config{}
	if err = yaml.Unmarshal(cfg, out); err != nil {
		return nil, err
	}
	return out, nil
}

