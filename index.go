package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func IndexGET(w http.ResponseWriter, r *http.Request, z httprouter.Params) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
func IndexPOST(w http.ResponseWriter, r *http.Request, z httprouter.Params) {
	id := DBInsert(r.FormValue("url"))
	encoded := strconv.FormatInt(id, 16)
	link := "http://localhost:8080/short/" + encoded

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "return.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	varmap := map[string]interface{}{
		"link": link,
	}
	if err := tmpl.ExecuteTemplate(w, "layout", varmap); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
