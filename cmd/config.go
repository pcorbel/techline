package main

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Config object to be loaded with configuration from YAML file.
type Config struct {
	DBHostname   string `yaml:"db_hostname,omitempty"`
	DBPort       int32  `yaml:"db_port,omitempty"`
	DBName       string `yaml:"db_name,omitempty"`
	DBUser       string `yaml:"db_user,omitempty"`
	DBPassword   string `yaml:"db_password,omitempty"`
	DBSSLMode    string `yaml:"db_ssl_mode,omitempty"`
	DBReset      bool   `yaml:"db_reset,omitempty"`
	LogLevel     string `yaml:"log_level,omitempty"`
	ServerPort   int32  `yaml:"server_port,omitempty"`
	WebDirectory string `yaml:"web_directory,omitempty"`
	Logger       *log.Logger
}

// LoadFromFile return Config object according to the YAML config file.
func (c *Config) LoadFromFile(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, c)
	}
	return err
}
