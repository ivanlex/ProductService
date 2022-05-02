package handlers

import (
	"ProductService/data"
	"context"
	"fmt"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
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
