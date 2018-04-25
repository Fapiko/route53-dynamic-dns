package main

import (
	"io/ioutil"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

/*
Config holds the configuration data this app requires to run
*/
type Config struct {
	Hostname           string `yaml:"hostname"`
	AwsAccessKeyID     string `yaml:"aws-access-key-id"`
	AwsSecretAccessKey string `yaml:"aws-secret-access-key"`
}

func parseConfig() (*Config, error) {
	usr, _ := user.Current()
	filepath := usr.HomeDir + string(os.PathSeparator) + ".route53-dynamic-dns/config.yaml"
	configData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configData, config)
	return config, err
}
