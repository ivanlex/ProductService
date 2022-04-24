package handlers

import (
	"ProductService/data"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(responseWriter http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle post method")

	Prod := request.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(Prod)

	p.l.Printf("Prod added: %#v", Prod)
}

func (p *Products) UpdateProducts(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle put method")

	Prod := request.Context().Value(KeyProduct{}).(*data.Product)
	data.UpdateProduct(Prod, id)

	p.l.Printf("Prod updated: %#v", Prod)
}

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

type KeyProduct struct {
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		Prod := &data.Product{}
		err := Prod.FromJson(request.Body)

		// unmarshal body to json
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// Validate the product
		err = Prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				responseWriter,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, Prod)
		req := request.WithContext(ctx)

		next.ServeHTTP(responseWriter, req)
	})
}
