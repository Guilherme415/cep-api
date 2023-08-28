package cmd

import (
	"fmt"
	"net/http"
)

func StartApi() {
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it's ok aroud here!")
}
