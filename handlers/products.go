package handlers

import (
	"encoding/json"
	"learn-golang/intro-to-microservices/product-api/data"
	"log"
	"net/http"
)

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

    rw.Write(d)
}
