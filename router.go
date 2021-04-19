package main

import (
	"fmt"
	"net/http"
)

type HandlersIface interface {
	Status(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func Router(controller HandlersIface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/status":
			switch r.Method {
			case "GET":
				controller.Status(w, r)
				return
			}
		case "/posts":
			switch r.Method {
			case "POST":
				controller.Create(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid route")
	}
}
