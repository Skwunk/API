package endpoints

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	w.WriteHeader(http.StatusOK)
}