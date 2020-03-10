package pkg

import "strings"

type vertex struct {
	value string
	label string
}

type graph interface {
	addEdge(v string, w string, label string)
	links(v string) []vertex
	toString() string
}

// Digraph - Graph impl
type Digraph struct {
	adj map[string]map[vertex]struct{}
}

var exists = struct{}{}

func newVertex(w string, label string) vertex {
	v := vertex{}
	v.value = w
	tLabel := strings.TrimSpace(label)
	if len(tLabel) > 0 {
		v.label = tLabel
	}
	return v
}

// AddEdge - edd edge to diagraph
func (d Digraph) AddEdge(v string, w string, label string) {
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

func (d Digraph) links(v string) map[vertex]struct{} {
	return d.adj[v]
}

// NewDigraph - Create a new Diagraph
func NewDigraph() *Digraph {
	return &Digraph{make(map[string]map[vertex]struct{})}
}

// ToString - Generate Diagraph dot file
func (d Digraph) ToString() string {
	var dString strings.Builder

	dString.WriteString("digraph {")
	for k, v := range d.adj {
		for w := range v {
			dString.WriteString(k)
			dString.WriteString(" -> ")
			dString.WriteString(w.value)
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
