package main

import "github.com/oliversavio/go-graph-generator/pkg"

func main() {
	g := pkg.NewDigraph()
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

	req := pkg.RenderRequest{Graph: g, DotFilePath: dotFilePath, GraphImagePath: outFilePath, ImageFormat: "-Tpng"}

	pkg.RenderGraph(req)

	s := pkg.GetSubgraph(g, "D")

	dotFilePath2 := "/tmp/hello2.dot"
	outFilePath2 := "/tmp/myImage2.png"

	req2 := pkg.RenderRequest{Graph: s, DotFilePath: dotFilePath2, GraphImagePath: outFilePath2, ImageFormat: "-Tpng"}

	pkg.RenderGraph(req2)

}
