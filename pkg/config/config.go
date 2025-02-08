package config

import "fmt"

type Config struct {
	KubeConfig  *KubeConfig  `yaml:"kube-config"`
	MongoConfig *MongoConfig `yaml:"mongo-config"`
}

type KubeConfig struct {
	Path string `yaml:"path"`
}

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (m *MongoConfig) Addr() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}
