package main

import (
	"fmt"

	"github.com/oliversavio/go-graph-generator/graphgen"
)

func main() {
	/* g := graphgen.NewDigraph()
	g.AddEdge("A", "B", "")
	g.AddEdge("A", "C", "")
	g.AddEdge("A", "E", "label 1")
	g.AddEdge("D", "C", "")
	g.AddEdge("B", "E", "")
	g.AddEdge("E", "F", "label 2")
	g.AddEdge("C", "F", "")
	g.AddEdge(1, "A", "ints")

	dotFilePath := "/tmp/hello.dot"
	outFilePath := "/tmp/myImage.png"

	req := graphgen.RenderRequest{Graph: g, DotFilePath: dotFilePath, GraphImagePath: outFilePath, ImageFormat: "-Tpng"}

	graphgen.RenderGraph(req)

	s := graphgen.GetSubgraph(g, "D")

	dotFilePath2 := "/tmp/hello2.dot"
	outFilePath2 := "/tmp/myImage2.png"

	req2 := graphgen.RenderRequest{Graph: s, DotFilePath: dotFilePath2, GraphImagePath: outFilePath2, ImageFormat: "-Tpng"}

	graphgen.RenderGraph(req2) */

	var graph graphgen.Graph
	digraph := graphgen.NewDigraph()
	graph = digraph

	graph.AddEdge("A", "B", "")
	graph.AddEdge("A", "C", "")
	graph.AddEdge("A", "E", "label 1")
	graph.AddEdge("D", "C", "")
	graph.AddEdge("B", "E", "")
	graph.AddEdge("E", "F", "label 2")
	graph.AddEdge("C", "F", "")
	fmt.Println(graph.ToString())

}
