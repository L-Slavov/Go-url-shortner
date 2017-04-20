package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)
	router.GET("/", IndexGET)
	router.POST("/", IndexPOST)
	router.GET("/short/:id", redirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
