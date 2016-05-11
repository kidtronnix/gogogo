# gogogo

fast and simple http routing. written in go.

## About

At it's core *gogogo* is a fast simple http router.
It matches requests to handlers by using a Trie data structure.
Typically this approach scales well.

As a bonus, *gogogo* will parse url params as it routes requests.

## Usage

```go
func id(w http.ResponseWriter, req *http.Request) {
  // get url param ':id'
  id := req.FormValue("id")
  // switch case for different methods, blah blah
}

func create(w http.ResponseWriter, req *http.Request) {
  // ...
}

func main() {

  // create router instance
	r := gogogo.NewRouter()

  // let's add a create handler for our resource
  r.Handle("/resource/", http.HandleFunc(create), "POST")

  // you can add multiple methods to a handler
	r.Handle("/resource/:id", http.HandleFunc(id), "GET", "PUT", "DELETE")

	http.ListenAndServe(":8000", r)
}

```
