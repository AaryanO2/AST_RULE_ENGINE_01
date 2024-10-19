package main

import (
	"fmt"
	"log"
	"net/http"
	"y/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
