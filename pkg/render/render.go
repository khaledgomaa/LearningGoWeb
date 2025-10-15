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

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("error creating template: ", err)
		}
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
