package graphgen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RenderRequest Params needed to render the graph
type RenderRequest struct {
	Graph          *Digraph
	DotFilePath    string
	GraphImagePath string
	ImageFormat    string
}

// RenderGraph - Generates the dot file and write it to a file
func RenderGraph(rr RenderRequest) {
	log.Println("Generating Graph dot file")
	writeToFile([]byte(rr.Graph.ToString()), rr.DotFilePath)
	graphBin := generateGraph(rr.DotFilePath, rr.ImageFormat)
	writeToFile(graphBin, rr.GraphImagePath)
	log.Println("Done")
}

func generateGraph(dotFilePath string, format string) []byte {
	cmd := exec.Command("dot", format, dotFilePath)
	o, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return o
}

func writeToFile(bytes []byte, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	n2, err := f.Write(bytes)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	log.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
