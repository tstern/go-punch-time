package main

import (
	"fmt"
	"os"

	"github.com/sternth/go-punch-time/routes"
)

var port string

func init() {
	var ok bool
	if port, ok = os.LookupEnv("PUNCH_TIME_PORT"); !ok {
		port = "8080"
	}
}

func main() {
	fmt.Printf("Start Server: http://localhost:%v/\n", port)
	routes.NewRouter(port)

}
