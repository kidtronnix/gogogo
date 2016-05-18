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
	if len(methods) == 0 {
		r.tree.addNode("mux", path, handler)
		return
	}
	for _, method := range methods {
		r.tree.addNode(method, path, handler)
	}
}

func (r *Router) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	r.Handle(path, http.HandlerFunc(handler), methods...)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], req)

	var handler http.Handler
	if node.muxHandler != nil {
		handler = node.muxHandler
	} else if h := node.methods[req.Method]; h != nil {
		handler = h
	} else if r.NotFoundHandler != nil {
		handler = r.NotFoundHandler
	} else {
		handler = http.HandlerFunc(http.NotFound)
	}
	handler.ServeHTTP(w, req)
}
