package db

import (
	"context"
	"github.com/zwtesttt/xzpCloud/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	opts := options.Client().ApplyURI("mongodb://192.168.3.63:27017").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	var result bson.M
	db = client.Database(databaseName)

	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return err
	}
	//fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}

func StopDatabase() error {
	return db.Client().Disconnect(context.TODO())
}
