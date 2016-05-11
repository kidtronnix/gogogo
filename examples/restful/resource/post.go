package resource

import (
	"fmt"
	"net/http"
)

func (e Endpoint) post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post")
}
