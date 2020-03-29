package graphgen

import "testing"

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
	graphTests("digraph {A -> B;}", g.ToString(), t)

	g.AddEdge("A", "B", "")
	graphTests("digraph {A -> B;}", g.ToString(), t)

	g.AddEdge("A", "C", "with edge")
	graphTests("digraph {A -> B;A -> C[label=\"with edge\"];}", g.ToString(), t)

}

func graphTests(expected string, actual string, t *testing.T) {
	if expected != actual {
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

	graphTests("digraph {C -> D;C -> E;}", g.GetSubgraph("C", 1).ToString(), t)

	graphTests("digraph {C -> D;C -> E;D -> F;D -> G;}", g.GetSubgraph("C", 2).ToString(), t)

	graphTests("digraph {E}", g.GetSubgraph("E", 2).ToString(), t)

}
