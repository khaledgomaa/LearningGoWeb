package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8090"
)

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response: ", n)
	fmt.Println("Headers: ", r.Header)
}

func main() {
	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(port, nil)
}
