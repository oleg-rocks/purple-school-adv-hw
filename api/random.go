package api

import (
	"math/rand"
	"net/http"
	"strconv"
)

type DiceNumberGenerator struct{}

func NewDiceNumberGenerator(router *http.ServeMux) {
	handler := &DiceNumberGenerator{}
	router.HandleFunc("/dice", handler.Generate())
}

func (d *DiceNumberGenerator) Generate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		res := strconv.Itoa(rand.Intn(6) + 1)
		w.Write([]byte(res))
	}
}
