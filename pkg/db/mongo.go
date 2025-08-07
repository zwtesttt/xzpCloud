package db

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zwtesttt/xzpCloud/pkg/config"
)

var (
	db *mongo.Database

	databaseName = "idc"
)

func GetDB() *mongo.Database {
	return db
}

func InitDatabase(cfg *config.Config) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// 构建连接字符串，支持认证
	var uri string
	if cfg.MongoConfig.Username != "" && cfg.MongoConfig.Password != "" {
		// URL编码用户名和密码以处理特殊字符
		username := url.QueryEscape(cfg.MongoConfig.Username)
		password := url.QueryEscape(cfg.MongoConfig.Password)
		uri = fmt.Sprintf("mongodb://%s:%s@%s",
			username,
			password,
			cfg.MongoConfig.Addr())
	} else {
		uri = fmt.Sprintf("mongodb://%s", cfg.MongoConfig.Addr())
	}

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	var result bson.M
	db = client.Database(databaseName)

	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return err
	}
	fmt.Println("Successfully connected to MongoDB!")
	return nil
}

func StopDatabase() error {
	return db.Client().Disconnect(context.TODO())
}
