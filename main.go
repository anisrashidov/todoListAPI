package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anisrashidov/todoAPP/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Welcome to your TODO list")
	fmt.Println("Listening on 127.0.0.1:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router.Router()))
}
