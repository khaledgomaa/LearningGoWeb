package main

import (
	"net/http"

	"github.com/learninggoweb/go-course/pkg/handlers"
)

const (
	port = ":8090"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	_ = http.ListenAndServe(port, nil)
}
