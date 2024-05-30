package main

import (
	"encoding/json"
	"fmt"
	"os"

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

	f, err := os.Create("./results/iterations.txt")
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}

	slice := make([]string, 1)

	randomChoser := disassemble.RandomChoser{ExcludeNodes: []int{}}
	ender := disassemble.NodeNumEnder{NodeNum: 3}
	iterationWriter := disassemble.StringSliceWriter{Slice: slice}

	disassemble.Disassemble(graphConfig, randomChoser, ender, iterationWriter)

	f.WriteString(iterationWriter.Slice[0])

	r, err := os.Create("./results/result.json")
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}
	byteStr, _ := json.Marshal(graphConfig)
	r.Write(byteStr)
	/*
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
		fmt.Printf("\n")*/
}
