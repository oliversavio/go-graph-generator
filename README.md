## GO Graph Generator

A simple utility to generate graphs using Graphviz

Currently only supports a simple Diagraph with Labels


Usage

{% highlight golang %}

// Create Graph
var g graphgen.Graph
digraph := graphgen.NewDigraph()
g = digraph

// Add Edges
g.AddEdge("A", "B", "")
g.AddEdge("A", "C", "")
// Add Edge with label
g.AddEdge("A", "E", "label 1")

// Add single Vertex
g.AddVertex("XXY")

{% endhighlight %}