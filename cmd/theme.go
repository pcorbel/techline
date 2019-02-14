package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Theme is a struct defining the type of an event.
type Theme struct {
	Title string `json:"title" yaml:"title"`
	Color string `json:"color" yaml:"color"`
	Icon  string `json:"icon" yaml:"icon"`
}

// Themes is a struct containing all themes.
type Themes struct {
	Themes []Theme `json:"themes" yaml:"themes"`
}

// LoadFromFile return Themes object according to the YAML config file.
func (s *Themes) LoadFromFile(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, s)
	}
	return err
}
