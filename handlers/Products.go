package handlers

import (
	"ProductService/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		p.GetProducts(responseWriter,request)
		return
	}

	responseWriter.WriteHeader(http.StatusNotFound)
}

func (p *Products)GetProducts(responseWriter http.ResponseWriter, request *http.Request)  {
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
