package resource

import "net/http"

// Handler is a struct that
type Endpoint struct{}

func (e Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		e.post(w, r)
	case "GET":
		e.get(w, r)
	case "PUT":
		// put specific methood middleware here...
		e.put(w, r)
	case "DELETE":
		e.delete(w, r)
	}
}
