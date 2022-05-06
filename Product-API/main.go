package main

import (
	"ProductService/handlers"
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	protos "github.com/kevin/currency/protos/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// make connection for grpc
	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// grpc client
	cc := protos.NewCurrencyClient(conn)

	productsHandler := handlers.NewProducts(logUtility, cc)

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

	// Middleware(from go-openapi/runtime) for Redoc
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)

	// Handle http files request and return swagger.yaml files
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Handel CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// Create http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
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
