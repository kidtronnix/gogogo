# gogogo

fast and simple http routing. written in go.

[![Build Status](https://travis-ci.org/smaxwellstewart/gogogo.svg?branch=master)](https://travis-ci.org/smaxwellstewart/gogogo)

## About


At it's core *gogogo* is a fast simple http router.
It matches requests to handlers by using a Trie data structure.
Typically this approach scales well.

As a bonus, *gogogo* will parse url params as it routes requests.

## Usage

```go
func createHandler(w http.ResponseWriter, req *http.Request) {
  // ...
}

func idHandler(w http.ResponseWriter, req *http.Request) {
  id := req.FormValue("id")
  // get :id url param then ...
}

func main() {

  // create router instance
  r := gogogo.NewRouter()

  // let's add a create handler for our resource
  r.HandleFunc("/resources/", createHanlder, "POST")

  // you can add multiple methods to a handler
  r.HandleFunc("/resources/:id", idHandler, "GET", "PUT", "DELETE")

  http.ListenAndServe(":8000", r)
}

```

# Credits

Great explanation of how trie data structure works in golang.

http://vluxe.io/golang-router.html
