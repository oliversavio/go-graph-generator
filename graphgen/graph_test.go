package graphgen

import (
	"sort"
	"strings"
	"testing"
)

func TestDiGraphImpl(t *testing.T) {
	var graph Graph
	digraph := NewDigraph()
	graph = digraph
	if graph == nil {
		t.Error("DiGraph was not initialized")
	}
}

func TestLinks(t *testing.T) {
	var g *Digraph = NewDigraph()
	g.AddEdge("A", "B", "")

	expected := "B"
	links := g.Links("A")

	for actual := range links {
		if expected != actual.Value {
			t.Errorf("Incorrect child %s for root %s", actual.Value, expected)
		}
	}

	noLinks := g.Links("B")
	if len(noLinks) > 0 {
		t.Error("Got invalid links for Diagraph")
	}

	g.AddEdge("A", "B", "")
	if len(g.Links("A")) != 1 {
		t.Errorf("Only 1 linked node expected, found >1 %s", g.ToString())
	}

	g.AddEdge("A", "B", "Hello")
	if len(g.Links("A")) != 2 {
		t.Errorf("Expected links are 2, found %d", len(g.Links("A")))
	}

}

func TestToString(t *testing.T) {
	var g *Digraph = NewDigraph()

	if "digraph {}" != g.ToString() {
		t.Error("Incorrect representation on empty diagraph")
	}

	g.AddEdge("A", "B", "")
	graphTests("digraph {A -> B;B;}", g.ToString(), t)

	g.AddEdge("A", "B", "")
	graphTests("digraph {A -> B;B;}", g.ToString(), t)

	g.AddEdge("A", "C", "with edge")
	graphTests("digraph {A -> B;A -> C[label=\"with edge\"];B;C;}", g.ToString(), t)

}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func graphTests(expected string, actual string, t *testing.T) {

	ex := sortString(expected)
	ac := sortString(actual)

	if ex != ac {
		t.Errorf("Expected: %s --> Actual: %s", expected, actual)
	}
}

func TestSubgraph(t *testing.T) {
	var g *Digraph = NewDigraph()

	g.AddEdge("A", "B", "")
	g.AddEdge("A", "C", "")
	g.AddEdge("C", "D", "")
	g.AddEdge("C", "E", "")
	g.AddEdge("D", "F", "")
	g.AddEdge("D", "G", "")

	graphTests("digraph {C -> D;C -> E;D;E;}", g.GetSubgraph("C", 1).ToString(), t)

	graphTests("digraph {C -> D;C -> E;D -> F;D -> G;E;F;G;}", g.GetSubgraph("C", 2).ToString(), t)

	graphTests("digraph {E;}", g.GetSubgraph("E", 2).ToString(), t)

}

func TestGetAllVertices(t *testing.T) {
	var g Graph
	digraph := NewDigraph()
	g = digraph

	g.AddEdge("A", "B", "none")
	vertices := g.V()

	g.AddEdge("A", "B", "")
	if len(vertices) != 2 {
		t.Errorf("Expected 2 Found: %d", len(vertices))
	}

	g.AddEdge("A", "C", "")
	g.AddEdge("C", "D", "")
	g.AddEdge("C", "E", "")
	g.AddEdge("D", "F", "")
	g.AddEdge("D", "G", "")

	subLevel1 := g.GetSubgraph("C", 10)

	if len(subLevel1.V()) != 5 {
		t.Errorf("Expected 5. Found: %d", len(subLevel1.V()))
	}

	subLevel2 := g.GetSubgraph("E", 10)
	if len(subLevel2.V()) != 1 {
		t.Errorf("Expected 1, Found %d", len(subLevel2.V()))
	}

}
