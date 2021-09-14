package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tstern/go-punch-time/route"
	"github.com/tstern/go-punch-time/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed loading .env file: %v", err)
	}

	port := flag.String("p", "8080", "port for webserver")
	flag.Parse()

	fmt.Println("connecting to database ...")
	uri := utils.CreateDbURI()
	db, err := utils.ConnectDb(uri)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}
	fmt.Println("connected to database")

	fmt.Println("starting server ...")
	err = route.NewRouter(*port, db)
	if err != nil {
		log.Fatalf("couldn't start server: %v", err)
	}
}
