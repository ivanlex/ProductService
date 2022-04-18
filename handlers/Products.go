package handlers

import (
	"ProductService/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		//Get
		p.GetProducts(responseWriter, request)
		return
	}

	if request.Method == http.MethodPut {
		//Put
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(request.URL.Path, -1)
		if len(g) != 1 {
			http.Error(responseWriter, "Invalid URI to update item", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(responseWriter, "Invalid Parameter to update item", http.StatusBadRequest)
			return
		}

		p.PutProducts(responseWriter, request, id)
	}

	if request.Method == http.MethodPost {
		//Post
		p.AddProducts(responseWriter, request)
	}

	if request.Method == http.MethodDelete {
		//Delete

		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(request.URL.Path, -1)
		if len(g) != 1 {
			http.Error(responseWriter, "Invalid URI to update item", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(responseWriter, "Invalid Parameter to update item", http.StatusBadRequest)
			return
		}

		p.DeleteProducts(responseWriter, request, id)
	}

	responseWriter.WriteHeader(http.StatusNotFound)
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

	Prod := &data.Product{}
	err := Prod.FromJson(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to add product from json", http.StatusBadRequest)
		return
	}

	data.AddProduct(Prod)

	p.l.Printf("Prod added: %#v", Prod)
}

func (p *Products) PutProducts(responseWriter http.ResponseWriter, request *http.Request, id int) {
	p.l.Println("Handle put method")

	Prod := &data.Product{}
	err := Prod.FromJson(request.Body)

	if err != nil {
		http.Error(responseWriter, "Unable to update product from json", http.StatusBadRequest)
		return
	}

	data.UpdateProduct(Prod, id)

	p.l.Printf("Prod updated: %#v", Prod)
}

func (p *Products) DeleteProducts(responseWriter http.ResponseWriter, request *http.Request, id int) {
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
