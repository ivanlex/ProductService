package handlers

import (
	"ProductService/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} productsResponse DeleteProducts
// Return a list of products from the database
// responses:
//	201: noContent

// DeleteProducts handles Delete a Product from database
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
