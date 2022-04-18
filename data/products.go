package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:idx`
	Name        string  `json:name`
	Description string  `json:description`
	Price       float32 `json:price`
	SKU         string  `json:sku`
	CreatedOn   string  `json:-`
	UpdatedOn   string  `json:-`
	DeletedOn   string  `json:-`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(writer io.Writer) error {
	e := json.NewEncoder(writer)
	return e.Encode(p)
}

func (p *Product) FromJson(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(p *Product, id int) {
	prod := getProductById(id)
	if prod != nil {
		prod.Name = p.Name
		prod.Price = p.Price
		prod.SKU = p.SKU
		prod.Description = p.Description
		prod.UpdatedOn = time.Now().String()
	} else {
		AddProduct(p)
	}
}

func DeleteProduct(id int) (Product, error) {
	prod := getProductById(id)
	if prod != nil {
		for index, item := range productList {
			if item.ID == id {
				productList = append(productList[:index], productList[index+1:]...)
				return *item, nil
			}
		}
	}

	return Product{}, errors.New("item not found")
}

func getProductById(id int) *Product {
	for _, item := range productList {
		if item.ID == id {
			return item
		}
	}

	return nil
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
