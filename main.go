package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	products := handlers.NewProducts(logger)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/products", products.GetProducts)
	putRouter.HandleFunc("/products/{id:[0-9]+}", products.UpdateProducts)
	putRouter.Use(products.MiddlewareProductValidation)
	postRouter.Use(products.MiddlewareProductValidation)
	postRouter.HandleFunc("/products", products.AddProduct)

	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Server is running and serving on port 3000")

		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Gracefully stopping server, command:", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), (30 * time.Second))
	server.Shutdown(timeoutContext)
}
