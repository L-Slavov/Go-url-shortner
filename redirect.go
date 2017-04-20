package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func redirect(w http.ResponseWriter, r *http.Request, z httprouter.Params) {
	idValue, err := strconv.Atoi(z.ByName("id"))
	if err != nil {
		fmt.Println("Couldn't parse ID")
	}
	result := DBget(idValue)
	http.Redirect(w, r, result, 301)
}
