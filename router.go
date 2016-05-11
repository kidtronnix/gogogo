package gogogo

import (
	"net/http"
	"strings"
)

func NewRouter() *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]http.Handler)}
	return &Router{tree: &node}
}

type Router struct {
	tree            *node
	NotFoundHandler http.Handler
}

func (r *Router) Handle(path string, handler http.Handler, methods ...string) {
	if path[0] != '/' {
		panic("Path has to start with a '/'.")
	}
	for _, method := range methods {
		r.tree.addNode(method, path, handler)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], req)

	if handler := node.methods[req.Method]; handler != nil {
		handler.ServeHTTP(w, req)
	} else if r.NotFoundHandler != nil {
		r.NotFoundHandler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}
