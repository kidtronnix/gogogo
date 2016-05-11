package resource

import (
	"fmt"
	"net/http"
)

func (e Endpoint) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete")
}
