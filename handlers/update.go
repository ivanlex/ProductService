package handlers

import (
	"ProductService/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (p *Products) UpdateProducts(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle put method")

	Prod := request.Context().Value(KeyProduct{}).(*data.Product)
	data.UpdateProduct(Prod, id)

	p.l.Printf("Prod updated: %#v", Prod)
}
