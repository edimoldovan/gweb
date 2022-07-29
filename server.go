package main

import (
	"fmt"
	"log"
	"main/utilities"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tmpl *template.Template

func main() {
	templateFiles := utilities.GetTemplates()
	tmpl, _ = template.ParseFiles(templateFiles...)
	router := httprouter.New()
	router.GET("/", Home)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "default", map[string]interface{}{
		"Title": "Web app with Go std",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
