package main

import (
	"log"
	"os"
)

var (
	port string
)

func init() {
	port = ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
}

func main() {
	r := NewRouter()

	log.Printf("APP is up, port(%v)", port)
	r.Run(port)
}