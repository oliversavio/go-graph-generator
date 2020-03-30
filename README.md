## GO Graph Generator

A simple utility to generate graphs using Graphviz

Currently only supports a simple `Diagraph` with Labels


### Usage

```golang
// Create Graph
var g graphgen.Graph
digraph := graphgen.NewDigraph()
g = digraph

// Add Edges
g.AddEdge("A", "B", "")
g.AddEdge("A", "C", "")
g.AddEdge("C", "F", "")
g.AddEdge("C", "E", "")

// Add Edge with label
g.AddEdge("A", "E", "label 1")

// Add single Vertex
g.AddVertex("XXY")

// Create a Subgraph with no of levels.
subGraph := g.GetSubgraph("C", 1)

// Render Graph as .png
// TODO: This needs a cleanup
dotFilePath := "/tmp/hello.dot"
outFilePath := "/tmp/myImage.png"
req := graphgen.RenderRequest{G: g, DotFilePath: dotFilePath, GraphImagePath: outFilePath, ImageFormat: "-Tpng"}
graphgen.RenderGraph(req)

```