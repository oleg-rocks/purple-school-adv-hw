package api

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type DiceNumberGenerator struct{
	Generator *rand.Rand
}

func NewDiceNumberGenerator(router *http.ServeMux) {
	handler := &DiceNumberGenerator{
		Generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	router.HandleFunc("/dice", handler.Generate())
}

func (d *DiceNumberGenerator) Generate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		res := strconv.Itoa(d.Generator.Intn(6) + 1)
		w.Write([]byte(res))
	}
}
