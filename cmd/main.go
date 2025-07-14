package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	server := http.Server {
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()
}