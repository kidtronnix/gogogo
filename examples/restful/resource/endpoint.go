package resource

import "net/http"

// Handler is a struct that
type Endpoint struct{}

func (e Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e.post(w, r)
	case http.MethodGet:
		e.get(w, r)
	case http.MethodPut:
		// put specific methood middleware here...
		e.put(w, r)
	case http.MethodDelete:
		e.delete(w, r)
	}
}
