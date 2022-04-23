package main

import (
	"ProductService/handlers"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Product service")
	//defer fmt.Println("Product Ended")
	logUtility := log.New(os.Stdout, "Product-API", log.LstdFlags)
	//playgroundHandler := handlers.NewPlayground(logUtility)
	productsHandler := handlers.NewProducts(logUtility)

	// Create a new Router by gorilla framework
	sm := mux.NewRouter()

	// SubRouter for each http methods
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	// Middleware handler for subRouter
	putRouter.Use(productsHandler.MiddlewareProductValidation)
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProducts)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	// Middleware handler for subRouter
	postRouter.Use(productsHandler.MiddlewareProductValidation)
	postRouter.HandleFunc("/products", productsHandler.AddProducts)

	// Create http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logUtility.Fatal(err)
		}
	}()

	// Listen system kill signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logUtility.Println("Received terminated, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
