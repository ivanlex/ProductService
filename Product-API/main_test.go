package main

import (
	"ProductService/client/client"
	"ProductService/client/client/products_response"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products_response.NewGetProductsParams()
	_, err := c.ProductsResponse.GetProducts(params)

	if err != nil {
		t.Fatal(err)
	}
}
