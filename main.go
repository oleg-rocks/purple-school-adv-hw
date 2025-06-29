package main

import (
	"fmt"
	"go/adv-hw/api"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	api.NewDiceNumberGenerator(router)
	server := http.Server {
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()
}