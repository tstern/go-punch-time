package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb(uri string) (*mongo.Database, error) {
	var dbName string
	var ok bool
	if dbName, ok = os.LookupEnv("MONGO_DB"); !ok {
		dbName = "punchTime"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return dbClient.Database(dbName), nil
}

func CreateDbURI() string {
	var username string
	var password string
	var host string
	var port string
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

	credentials := ""
	if username != "" {
		if password != "" {
			credentials = fmt.Sprintf("%v:%v@", username, password)
		} else {
			credentials = fmt.Sprintf("%v@", username)
		}
	}

	return fmt.Sprintf("mongodb://%v%v:%v", credentials, host, port)
}
