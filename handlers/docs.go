// Package classification Products API.
//
// Documentation for Product API
//
// Terms Of Service:
//
// Handler for products
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handlers

import "ProductService/data"

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All current products
	// in: body
	Body []data.Product
}

// empty response
// swagger:response noContent
type noContentWrapper struct {
}

// swagger:parameters DeleteProducts
type productIDParameter struct {
	// The id of product to delete from the database
	// in : path
	// required: true
	ID int `json:"id"`
}
