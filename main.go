package main

import (
	"fmt"
	"go-test-supporting-project/api"
	"go-test-supporting-project/controllers"
	"log"
	"net/http"
)

const APP_PORT = 8090

func main() {
	handlers := controllers.Handler{
		Store: api.ApiRequester{BaseUrl: "https://jsonplaceholder.typicode.com"},
	}
	router := Router(handlers)

	fmt.Printf("Starting server at 0.0.0.0:%d\n", APP_PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", APP_PORT), http.HandlerFunc(router))

	if err != nil {
		log.Fatal(err)
	}
}
