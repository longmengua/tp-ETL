package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Init(rootPath string) *Config {
	c := &Config{}
	filePath := fmt.Sprintf("%s/conf.yaml", rootPath)
	log.Printf("Loading env file: %s", filePath)

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("failed to init config ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Panic("failed to unmarshal config", err)
	}
	return c
}
