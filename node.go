package gogogo

import (
	"net/http"
	"strings"
)

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	muxHandler   http.Handler
	methods      map[string]http.Handler
}

// addNode - adds a node to our tree. Will add multiple nodes if path
// can be broken up into multiple components. Those nodes will have no
// handler implemented and will fall through to the default handler.
func (n *node) addNode(method, path string, handler http.Handler) {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	for {

		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 { // update an existing node.
			if method == "mux" {
				aNode.muxHandler = handler
			} else {
				aNode.methods[method] = handler
			}
			return
		}
		newNode := node{component: component, isNamedParam: false, methods: make(map[string]http.Handler)}

		if len(component) > 0 && component[0] == ':' { // check if it is a named param.
			newNode.isNamedParam = true
		}
		if count == 1 { // this is the last component of the url resource, so it gets the handler.
			if method == "mux" {
				newNode.muxHandler = handler
			} else {
				newNode.methods[method] = handler
			}

		}
		aNode.children = append(aNode.children, &newNode)

		count--
		if count == 0 {
			break
		}
	}
}

// traverse scans over a slice of components to
func (n *node) traverse(components []string, req *http.Request) (*node, string) {
	component := components[0]
	if len(n.children) > 0 { // no children, then bail out.
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && req != nil {
					req.Form.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 { // http://xkcd.com/1270/
					return child.traverse(next, req) // tail recursion is it's own reward.
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
