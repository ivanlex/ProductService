package handlers

import (
	"ProductService/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (p *Products) DeleteProducts(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle delete method")

	Prod, err := data.DeleteProduct(id)

	if err != nil {
		p.l.Printf("Prod deleted item can't found", Prod)
		http.Error(responseWriter, "Delete item can't found", http.StatusBadRequest)
		return
	} else {
		p.l.Printf("Prod deleted: %#v", Prod)
	}
}
