package db

import (
	"testing"

	"github.com/zwtesttt/xzpCloud/pkg/config"
)

func TestMongo(t *testing.T) {
	err := InitDatabase(&config.Config{
		MongoConfig: &config.MongoConfig{
			Host:     "127.0.0.1",
			Port:     27017,
			Username: "root",
			Password: "qwe@123",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
}
