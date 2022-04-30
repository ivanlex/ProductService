package handlers

import (
	"ProductService/data"
	"net/http"
)

// swagger:route GET /products productsResponse GetProducts
// Return a list of products from the database
// responses:
//	200: productResponse

// GetProducts handles GET requests
func (p *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
