package handlers

import (
	"fmt"
	"github.com/learninggoweb/go-course/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func OldHome(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response: ", n)
	fmt.Println("Headers: ", r.Header)
}
