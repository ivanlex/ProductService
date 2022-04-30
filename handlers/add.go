package handlers

import (
	"ProductService/data"
	"net/http"
)

func (p *Products) AddProducts(responseWriter http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle post method")

	Prod := request.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(Prod)

	p.l.Printf("Prod added: %#v", Prod)
}
