package graphgen

import (
	"container/list"
	"strconv"
	"strings"
)

// Vertex - Hold vertex data
type Vertex struct {
	Value interface{}
	Label string
}

// Graph - Graph DS
type Graph interface {
	AddVertex(v interface{}) *Vertex
	AddEdge(v interface{}, w interface{}, label string)
	Links(v interface{}) map[*Vertex]struct{}
	ToString() string
	GetSubgraph(root interface{}, maxLevel int) Graph
	V() []Vertex
}

// Digraph - Graph impl. Supports ints, string and bool types
type Digraph struct {
	st  map[interface{}]*Vertex
	adj map[*Vertex]map[*Vertex]struct{}
}

var exists = struct{}{}

func newVertex(w interface{}, label string) *Vertex {
	v := Vertex{}
	v.Value = w
	tLabel := strings.TrimSpace(label)
	if len(tLabel) > 0 {
		v.Label = tLabel
	}
	return &v
}

// AddEdge - edd edge to diagraph
func (d *Digraph) AddEdge(v interface{}, w interface{}, label string) {
	var vert *Vertex = d.AddVertex(v)
	var wert *Vertex = d.AddVertex(w)
	edges, _ := d.adj[vert]
	if wert.Label != label {
		edges[newVertex(w, label)] = exists
	} else {
		edges[wert] = exists
	}

	d.adj[vert] = edges
}

// Links - Get linked vertices
func (d *Digraph) Links(v interface{}) map[*Vertex]struct{} {
	vtx, ok := d.st[v]
	if ok {
		return d.adj[vtx]
	}
	return make(map[*Vertex]struct{})
}

// NewDigraph - Create a new Diagraph
// Currently Diagraph supports ints, bool and string types
func NewDigraph() *Digraph {
	return &Digraph{make(map[interface{}]*Vertex), make(map[*Vertex]map[*Vertex]struct{})}
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

// GetSubgraph - Crates a new graph from the node specified as root
func (d *Digraph) GetSubgraph(root interface{}, maxLevel int) Graph {
	var sub Graph
	digraph := NewDigraph()
	sub = digraph
	q := list.New()
	q.PushBack(d.st[root])

	currentLevel := 0

	for q.Len() > 0 && currentLevel < maxLevel {
		currentLevel++
		current := q.Front()
		v := q.Remove(current).(*Vertex)
		sub = d.populateChildren(sub, v.Value)
		for ch := range d.Links(v.Value) {
			q.PushBack(ch)
		}
	}

	return sub
}

func (d *Digraph) populateChildren(sg Graph, parent interface{}) Graph {
	children := d.Links(parent)
	for k := range children {
		sg.AddEdge(getString(parent), getString(k.Value), k.Label)
	}

	if len(children) == 0 {
		sg.AddVertex(parent)
	}

	return sg
}

// V - Get all vertices of graph
func (d *Digraph) V() []Vertex {
	vertices := make([]Vertex, 0)
	for v := range d.adj {
		vertices = append(vertices, *v)
	}
	return vertices
}

// AddVertex - Add vertext to Graph
func (d *Digraph) AddVertex(v interface{}) *Vertex {
	vtx, exist := d.st[v]
	if !exist {
		vtx = newVertex(v, "")
		d.st[v] = vtx
		d.adj[vtx] = make(map[*Vertex]struct{})
	}
	return vtx
}

// ToString - Generate Diagraph dot file
func (d *Digraph) ToString() string {
	var dString strings.Builder
	dString.WriteString("digraph {")
	for _, v := range d.V() {
		children := d.Links(v.Value)
		if len(children) > 0 {
			renderChildren(&dString, v, children)
		} else {
			dString.WriteString(getString(v.Value))
			dString.WriteString(";")
		}
	}
	dString.WriteString("}")

	return dString.String()
}

func renderChildren(dString *strings.Builder, root Vertex, children map[*Vertex]struct{}) *strings.Builder {
	for w := range children {
		dString.WriteString(getString(root.Value))
		dString.WriteString(" -> ")
		dString.WriteString(getString(w.Value))
		if len(w.Label) > 0 {
			dString.WriteString("[label=\"")
			dString.WriteString(w.Label)
			dString.WriteString("\"]")
		}
		dString.WriteString(";")
	}

	return dString
}
