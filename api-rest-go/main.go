package main

import (
	"html/template"
	"log"
	"net/http"
)

// Context is the page context
type Context struct {
	Fecha        string `json:"date"`
	DolarOficial string `json:"officialdolar"`
	DolarBlue    string `json:"bluedollar"`
	Variacion    string `json:"variation"`
}

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			ctx := Context{Fecha: "mandale el parametro"}
			err := t.Execute(w, ctx)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
