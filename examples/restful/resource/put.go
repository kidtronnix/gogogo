package resource

import (
	"fmt"
	"net/http"
)

func (e Endpoint) put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "put")
}
