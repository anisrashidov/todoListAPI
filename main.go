package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anisrashidov/todoAPP/router"
)

func main() {
	fmt.Println("Welcome to your TODO list")
	log.Fatal(http.ListenAndServe(":8000", router.Router()))
	fmt.Println("Listening on 127.0.0.1:8000")
}
