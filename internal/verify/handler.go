package validation

import (
	"fmt"
	"go/adv-hw/configs"
	"net/http"
)

type ValidationHandlerDeps struct {
	*configs.Config
}

type ValidationHandler struct {
	*configs.Config
}

func NewValidationHandler(router *http.ServeMux, deps ValidationHandlerDeps) {
	handler := &ValidationHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("POST /verify", handler.Verify())
}	

func (handler *ValidationHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Send")
	}
}

func (handler *ValidationHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Verify")
	}
}
