package handlers

import (
	"learn-golang/intro-to-microservices/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect the product id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "invalid uri", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "invalid uri", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			http.Error(rw, "invalid uri", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch products from datastore
	lp := data.GetProducts()

	// serialize list to json
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)

	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		} else {
			http.Error(rw, "Unable to update product", http.StatusInternalServerError)
			return
		}
	}
}
