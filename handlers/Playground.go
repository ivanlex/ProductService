package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Playground struct {
	l *log.Logger
}

func NewPlayground(l *log.Logger) *Playground {
	return &Playground{l}
}

func (p *Playground) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	p.l.Println("Product API")

	data, err := ioutil.ReadAll(request.Body)

	if err != nil {
		p.l.Println("Error found")
		http.Error(responseWriter, "Error found", http.StatusBadRequest)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(fmt.Sprintf("Input data is %s", data)))
}
