package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var instance *mongo.Database
var username string
var password string
var host string
var port string
var db string

func init() {
	var ok bool
	if username, ok = os.LookupEnv("MONGO_INITDB_ROOT_USERNAME"); !ok {
		username = ""
	}
	if password, ok = os.LookupEnv("MONGO_INITDB_ROOT_PASSWORD"); !ok {
		password = ""
	}
	if host, ok = os.LookupEnv("MONGO_HOST"); !ok {
		host = "localhost"
	}
	if port, ok = os.LookupEnv("MONGO_PORT"); !ok {
		port = "27017"
	}
	if db, ok = os.LookupEnv("MONGO_DB"); !ok {
		db = "punchTime"
	}
}

func ConnectDb() *mongo.Database {
	once.Do(func() {
		uri := createUri()
		fmt.Printf("Connecting to database %v ...\n", uri)
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		instance = client.Database(db)
		fmt.Printf("Connected to database \"%v\"\n", instance.Name())
	})
	return instance
}

func createUri() string {
	credentials := ""
	if username != "" && password != "" {
		credentials = fmt.Sprintf("%v:%v@", username, password)
	}
	return fmt.Sprintf("mongodb://%v%v:%v", credentials, host, port)

}
