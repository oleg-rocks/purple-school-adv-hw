package verify

import (
	"fmt"
	"go/adv-hw/configs"
	"go/adv-hw/pkg/gen"
	"go/adv-hw/pkg/req"
	"go/adv-hw/pkg/res"
	"go/adv-hw/pkg/sender"
	"go/adv-hw/pkg/store"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router chi.Router, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}	

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
		}
		hash := gen.GenerateHash(payload.Email)
		data := store.VerifyData{
			Email: payload.Email,
			Hash: hash,
		}
		store.StoreVerifyData(data)
		fmt.Println(data)
		err = sender.SendVerification(payload.Email, hash, handler.Config)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")
		match, err := store.HasHash(hash)
		if err != nil {
			fmt.Println(err.Error())
		}
		if match {
			fmt.Println("✅ Email is verified!")
		} else {
			fmt.Println("❌ Failed to verify this email.")
		}
	}
}
