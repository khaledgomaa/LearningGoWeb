package main

import (
	"net/http"
)

const (
	port = ":8090"
)

func main() {
	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(port, nil)
}
