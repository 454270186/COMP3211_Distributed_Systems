package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
	 port = ":" + val
	}
	return port
   }

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func greet2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World22222")
}

func main() {
	http.HandleFunc("/api/hello", greet)
	http.HandleFunc("/api/hello2", greet2)
	http.ListenAndServe(get_port(), nil)
}