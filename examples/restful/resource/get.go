package resource

import (
	"fmt"
	"net/http"
)

func (e Endpoint) get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	switch id {
	case "":
		index(w, r)
	default:
		getOne(w, r)
	}
}

func getOne(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Fprintf(w, "GET: %s", id)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET: index of all resources")
}
