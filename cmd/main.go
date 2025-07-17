package main

import (
	"fmt"
	"go/adv-hw/configs"
	"go/adv-hw/internal/verify"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	config := configs.LoadConfig()
	router := chi.NewRouter()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: config,
	})
	server := http.Server {
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()
}