package controllers

import (
	"fmt"
	"net/http"
)

func (h Handler) Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status OK")
}
