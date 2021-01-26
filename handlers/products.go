package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
