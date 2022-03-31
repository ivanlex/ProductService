package main

import (
	"ProductService/handlers"
	"context"
	"fmt"
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
	playgroundHandler := handlers.NewPlayground(logUtility)
	productsHandler := handlers.NewProducts(logUtility)

	sm := http.NewServeMux()
	sm.Handle("/", playgroundHandler)
	sm.Handle("/products", productsHandler)

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

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logUtility.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
