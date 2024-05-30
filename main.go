package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/FackOff25/disassemble_assemble/assemble"
	"github.com/FackOff25/disassemble_assemble/disassemble"
	"github.com/FackOff25/disassemble_assemble/graph"
	"github.com/FackOff25/disassemble_assemble/microsolution"
)

func writeFloatMatrix(matrix [][]float64, file string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	for i := range matrix {
		str := fmt.Sprintf("%g", matrix[i][0])
		for _, v := range matrix[i][1:] {
			str = fmt.Sprintf("%s\t%g", str, v)
		}
		str += "\n"
		f.WriteString(str)
	}
	f.Close()
}

func writeIntMatrix(matrix [][]int, file string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	for i := range matrix {
		str := fmt.Sprintf("%d", matrix[i][0])
		for _, v := range matrix[i][1:] {
			str = fmt.Sprintf("%s\t%d", str, v)
		}
		str += "\n"
		f.WriteString(str)
	}
	f.Close()
}

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

	idToNum := disassemble.MakeIdToNum(graphConfig)
	fmt.Println(idToNum)

	slice := make([]string, 1)

	randomChoser := disassemble.RandomChoser{ExcludeNodes: []int{1, 2}}
	ender := disassemble.NodeNumEnder{NodeNum: 2}
	iterationWriter := disassemble.StringSliceWriter{Slice: slice}

	disassemble.Disassemble(graphConfig, randomChoser, ender, iterationWriter)

	f.WriteString(iterationWriter.Slice[0])
	f.Close()

	r, err := os.Create("./results/result.json")
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}
	byteStr, _ := json.Marshal(graphConfig)
	r.Write(byteStr)

	M, P := microsolution.SolveOnePath(graphConfig, idToNum, 1, 2)

	f, err = os.Open("./results/iterations.txt")
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}
	iterationReader := assemble.StraightScannerIterationReader{Reader: bufio.NewReader(f)}

	remainNodes := make([]int, len(graphConfig.Nodes))
	counter := 0
	for k := range graphConfig.Nodes {
		remainNodes[counter] = k
		counter++
	}

	assemble.Assemble(M, P, remainNodes, idToNum, iterationReader)

	writeFloatMatrix(M, "./results/M.txt")
	writeIntMatrix(P, "./results/P.txt")
	//fmt.Printf("M: \n %v\n", M)
	//fmt.Printf("P: \n %v\n", P)
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
