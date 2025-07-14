package main

import (
	"fmt"
	"go/adv-hw/configs"
	"go/adv-hw/internal/verify"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	server := http.Server {
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: config,
	})
}