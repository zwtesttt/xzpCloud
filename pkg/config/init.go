package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Init(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer file.Close()

	// 解码 YAML 文件内容
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}

	return &config
}
