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

	dotFilePath := "/tmp/hello.dot"
	outFilePath := "/tmp/myImage.png"

	req := pkg.RenderRequest{Graph: g, DotFilePath: dotFilePath, GraphImagePath: outFilePath, ImageFormat: "-Tpng"}

	pkg.RenderGraph(req)

}
