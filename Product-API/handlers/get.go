package handlers

import (
	"ProductService/data"
	"context"
	protos "github.com/kevin/currency/protos/currency"
	"net/http"
)

// swagger:route GET /products productsResponse GetProducts
// Return a list of products from the database
// responses:
//	200: productResponse

// GetProducts handles GET requests
func (p *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	//Add response type to client,or you will get failed test this function.
	responseWriter.Header().Add("Content-Type", "application/json")
	listProducts := data.GetProducts()

	//request by grpc
	rr := &protos.RateRequest{
		Base:        protos.Currencies_GBP,
		Destination: protos.Currencies_JPY,
	}
	response, err := p.cc.GetRate(context.Background(), rr)

	p.l.Printf("Rate is %v", response.Rate)

	prod2 := make([]*data.Product, len(listProducts))

	// make new instance for listProducts
	for index, item := range listProducts {
		prod2[index] = &data.Product{
			ID:          item.ID,
			Name:        item.Name,
			SKU:         item.SKU,
			Price:       item.Price * response.Rate,
			Description: item.Description,
		}
	}

	items := data.Products(prod2)
	items.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
