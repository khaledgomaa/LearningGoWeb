package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tml, "./templates/base.layout.tmpl") // Added template file while rendering the templates
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
