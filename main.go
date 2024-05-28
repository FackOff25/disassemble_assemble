package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FackOff25/disassemble_assemble/astar"
	"github.com/FackOff25/disassemble_assemble/disassemble"
	"github.com/FackOff25/disassemble_assemble/graph"
)

func main() {
	graphConfig, err := graph.Read("./testData/testFile.json")
	//fmt.Printf("%#v\n", graphConfig)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}

	f, err := os.Create("./results/result.txt")
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}

	randomChoser := disassemble.RandomChoser{ExcludeNodes: []int{}}
	ender := disassemble.NodeNumEnder{NodeNum: 2}
	iterationWriter := disassemble.WriterIterationWriter{Writer: *bufio.NewWriter(f)}

	disassemble.Disassemble(graphConfig, randomChoser, ender, iterationWriter)

	path, dist, found := astar.Path(graphConfig.Nodes[31], graphConfig.Nodes[25])

	if !found {
		fmt.Printf("No path\n")
		return
	}

	fmt.Printf("Path: dist=%f\n", dist)
	fmt.Printf("%d", path[0].GetId())
	for _, v := range path[1:] {
		fmt.Printf("->%d", v.GetId())
	}
	fmt.Printf("\n")
}
