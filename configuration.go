package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Application App `yaml:"app"`
}

func loadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}
