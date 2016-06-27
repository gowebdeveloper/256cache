package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/",mapdoing)
	// Bind to a port and pass our router in
	http.ListenAndServe(":8000", nil)
}