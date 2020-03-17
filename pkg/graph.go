package pkg

import (
	"strconv"
	"strings"
)

type vertex struct {
	value interface{}
	label string
}

type graph interface {
	addEdge(v interface{}, w interface{}, label string)
	links(v interface{}) []vertex
	toString() string
}

// Digraph - Graph impl. Supports ints, string and bool types
type Digraph struct {
	adj map[interface{}]map[vertex]struct{}
}

var exists = struct{}{}

func newVertex(w interface{}, label string) vertex {
	v := vertex{}
	v.value = w
	tLabel := strings.TrimSpace(label)
	if len(tLabel) > 0 {
		v.label = tLabel
	}
	return v
}

// AddEdge - edd edge to diagraph
func (d Digraph) AddEdge(v interface{}, w interface{}, label string) {
	edges, exist := d.adj[v]
	if exist {
		edges[newVertex(w, label)] = exists
		d.adj[v] = edges
	} else {
		newEdges := make(map[vertex]struct{})
		newEdges[newVertex(w, label)] = exists
		d.adj[v] = newEdges
	}
}

func (d Digraph) links(v interface{}) map[vertex]struct{} {
	return d.adj[v]
}

// NewDigraph - Create a new Diagraph
// Currently Diagraph supports ints, bool and string types
func NewDigraph() *Digraph {
	return &Digraph{make(map[interface{}]map[vertex]struct{})}
}

func getString(i interface{}) string {

	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	default:
		return "N/A"
	}
}

// ToString - Generate Diagraph dot file
func (d Digraph) ToString() string {
	var dString strings.Builder

	dString.WriteString("digraph {")
	for k, v := range d.adj {
		for w := range v {
			dString.WriteString(getString(k))
			dString.WriteString(" -> ")
			dString.WriteString(getString(w.value))
			if len(w.label) > 0 {
				dString.WriteString("[label=\"")
				dString.WriteString(w.label)
				dString.WriteString("\"]")
			}
			dString.WriteString(";")
		}
	}
	dString.WriteString("}")

	return dString.String()
}
