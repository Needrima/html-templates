package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("editors.html"))
}

func main() {

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			if err := tpl.Execute(w, nil); err != nil {
				log.Fatal("Error:", err)
			}

		case http.MethodPost:
			data := r.FormValue("description")
			if err := tpl.Execute(w, data); err != nil {
				log.Fatal("Error:", err)
			}

		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
