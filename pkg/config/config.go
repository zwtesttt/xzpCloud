package config

import (
	"fmt"
	"time"
)

type Config struct {
	KubeConfig  *KubeConfig  `yaml:"kube-config"`
	MongoConfig *MongoConfig `yaml:"mongo-config"`
	Log         *LogConfig   `yaml:"log"`
}

type KubeConfig struct {
	Path string `yaml:"path"`
}

type LogConfig struct {
	Level           string        `yaml:"level"`
	SkipPaths       []string      `yaml:"skip-paths"`
	LogRequestBody  bool          `yaml:"log-request-body"`
	LogResponseBody bool          `yaml:"log-response-body"`
	MaxBodySize     int64         `yaml:"max-body-size"`
	SlowThreshold   time.Duration `yaml:"slow-threshold"`
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
