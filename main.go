package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sternth/go-punch-time/route"
	"github.com/sternth/go-punch-time/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed loading .env file: %v", err)
	}

	port := flag.String("p", "8080", "port for webserver")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uri := utils.CreateDbURI()
	db, err := utils.ConnectDb(ctx, uri)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	err = route.NewRouter(*port, db)
	if err != nil {
		log.Fatalf("couldn't start server: %v", err)
	}
}
