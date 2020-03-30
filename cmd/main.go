package main

import (
	"github.com/oliversavio/go-graph-generator/graphgen"
)

func main() {
	var g graphgen.Graph
	digraph := graphgen.NewDigraph()

	g = digraph
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

	req := graphgen.RenderRequest{G: g, DotFilePath: dotFilePath, GraphImagePath: outFilePath, ImageFormat: "-Tpng"}

	graphgen.RenderGraph(req)

	s := g.GetSubgraph("D", 10)

	dotFilePath2 := "/tmp/hello2.dot"
	outFilePath2 := "/tmp/myImage2.png"

	req2 := graphgen.RenderRequest{G: s, DotFilePath: dotFilePath2, GraphImagePath: outFilePath2, ImageFormat: "-Tpng"}

	graphgen.RenderGraph(req2)

}
