package main

import (
	"net/http"

	"github.com/smaxwellstewart/gogogo"
	"github.com/smaxwellstewart/gogogo/examples/restful/resource"
)

func main() {
	r := gogogo.NewRouter()

	r.Handle("/resources/", resource.Endpoint{}, "POST")
	r.Handle("/resources/:id", resource.Endpoint{}, "GET", "PUT", "DELETE")

	http.ListenAndServe(":8000", r)
}
