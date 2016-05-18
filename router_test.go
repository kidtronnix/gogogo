package gogogo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update %s", r.Method)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "notfound")
}

func TestNewRouter(t *testing.T) {
	assert := assert.New(t)

	r := NewRouter()

	expectedNewRouter := &Router{
		tree: &node{
			component:    "/",
			isNamedParam: false,
			methods:      make(map[string]http.Handler),
		},
	}
	assert.Equal(expectedNewRouter, r)
}

func TestRouterCorrectlyMuxes(t *testing.T) {
	assert := assert.New(t)

	// setup our router
	r := NewRouter()
	r.HandleFunc("/resources/", index)
	r.Handle("/resources/:id", http.HandlerFunc(update))

	// test index
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/resources/", nil)
	r.ServeHTTP(w, req)
	assert.Equal("index", w.Body.String())

	// test id route
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "http://example.com/resources/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal("update PUT", w.Body.String())
}

func TestRouterCorrectlyRoutesMethods(t *testing.T) {
	assert := assert.New(t)

	// setup our router
	r := NewRouter()
	r.Handle("/", http.HandlerFunc(root), "GET")
	r.HandleFunc("/resources/", index, "GET")
	r.Handle("/resources/:id", http.HandlerFunc(update), "PUT", "PATCH")

	// test root
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	r.ServeHTTP(w, req)
	assert.Equal("root", w.Body.String())

	// test index
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "http://example.com/resources/", nil)
	r.ServeHTTP(w, req)
	assert.Equal("index", w.Body.String())

	// test id route
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "http://example.com/resources/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal("update PUT", w.Body.String())
}

func TestRouterParsesURLParams(t *testing.T) {
	assert := assert.New(t)

	r := NewRouter()
	r.Handle("/:id", http.HandlerFunc(root), "GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal("1", req.FormValue("id"))
}

func TestHandlePanicsOnBadPath(t *testing.T) {
	assert := assert.New(t)

	r := NewRouter()

	assert.Panics(func() {
		r.Handle("badpath", http.HandlerFunc(root), "GET")
	})
}

func TestHandleNotFoundHandler(t *testing.T) {
	assert := assert.New(t)

	r := NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notfound)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "http://example.com/", nil)
	r.ServeHTTP(w, req)
	assert.Equal("notfound", w.Body.String())
}

func TestHandleDefaultNotFoundHandler(t *testing.T) {
	assert := assert.New(t)

	r := NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "http://example.com/", nil)
	r.ServeHTTP(w, req)

	_w := httptest.NewRecorder()
	http.NotFound(_w, req)

	assert.Equal(_w, w)
}
