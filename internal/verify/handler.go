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
			return
		}
		hash := gen.GenerateHash(payload.Email)
		data := store.VerifyData{
			Email: payload.Email,
			Hash: hash,
		}
		store.Save(data)
		fmt.Println(data)
		err = sender.SendVerification(payload.Email, hash, handler.Config)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")
		match, err := store.FindHashAndRemove(hash)
		if err != nil {
			fmt.Println(err.Error())
			res.Json(w, false, http.StatusNotFound)
			return
		}
		if match {
			fmt.Println("✅ Email is verified!")
			res.Json(w, true, http.StatusOK)
		} else {
			fmt.Println("❌ Failed to verify this email.")
			res.Json(w, false, http.StatusNotFound)
		}
	}
}
