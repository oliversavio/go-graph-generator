package graphgen

import (
	"container/list"
	"log"
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
	vtx, vExist := d.st[v]
	wtx, wExist := d.st[w]

	if !vExist {
		log.Printf("Vertex %s doesn't exist, creating", v)
		vtx = newVertex(v, "")
		d.st[v] = vtx
	}

	if !wExist || wtx.Label != label {
		wtx = newVertex(w, label)
		d.st[w] = wtx
	}

	edges, ok := d.adj[vtx]
	if ok {
		edges[wtx] = exists
		d.adj[vtx] = edges
	} else {
		newEdges := make(map[*Vertex]struct{})
		newEdges[wtx] = exists
		d.adj[vtx] = newEdges
		d.adj[wtx] = make(map[*Vertex]struct{})
	}
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
	return sg
}

// V - Get all vertices of graph
func (d *Digraph) V() []Vertex {
	vertices := make([]Vertex, len(d.adj))
	for v := range d.adj {
		vertices = append(vertices, *v)
	}
	return vertices
}

// ToString - Generate Diagraph dot file
func (d *Digraph) ToString() string {
	var dString strings.Builder
	// TODO: Needs fixing
	dString.WriteString("digraph {")
	for k, v := range d.adj {
		for w := range v {
			dString.WriteString(getString(k.Value))
			dString.WriteString(" -> ")
			dString.WriteString(getString(w.Value))
			if len(w.Label) > 0 {
				dString.WriteString("[label=\"")
				dString.WriteString(w.Label)
				dString.WriteString("\"]")
			}
			dString.WriteString(";")
		}
	}
	dString.WriteString("}")

	return dString.String()
}
