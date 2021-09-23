package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")

}

// Checkout out http://localhost:8000/ after running
