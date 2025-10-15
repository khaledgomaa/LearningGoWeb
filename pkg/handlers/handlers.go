package handlers

import (
	"fmt"
	"net/http"

	"github.com/learninggoweb/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateTest(w, "home.page.tmpl")
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
