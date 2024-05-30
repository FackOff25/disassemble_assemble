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

const (
	testFile      = "./testData/testFile.json"
	iterationFile = "./results/iterations.txt"
	resultFile    = "./results/result.json"
	idToNumFile   = "./results/idToNum.txt"
	mMatrixFile   = "./results/M.txt"
	pMatrixFile   = "./results/P.txt"
)

func errorCheck(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

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
	errorCheck(err)

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

func makeIdToNum(graphConfig graph.Graph, idToNumFile string) disassemble.IdToNum {
	idToNum := disassemble.MakeIdToNum(graphConfig)
	m, err := os.Create(idToNumFile)
	errorCheck(err)
	m.WriteString(fmt.Sprintf("%v", idToNum))
	m.Close()
	return idToNum
}

func doDisassemble(graphConfig graph.Graph, iterationFile, resultfile string) {
	f, err := os.Create(iterationFile)
	errorCheck(err)

	slice := make([]string, 1)

	randomChoser := disassemble.RandomChoser{ExcludeNodes: []int{1, 2}}
	ender := disassemble.NodeNumEnder{NodeNum: 3}
	iterationWriter := disassemble.StringSliceWriter{Slice: slice}

	disassemble.Disassemble(graphConfig, randomChoser, ender, iterationWriter)

	f.WriteString(iterationWriter.Slice[0])
	f.Close()

	r, err := os.Create(resultfile)
	errorCheck(err)
	byteStr, _ := json.Marshal(graphConfig)
	r.Write(byteStr)
}

func doAssemble(M [][]float64, P [][]int, graphConfig graph.Graph, idToNum disassemble.IdToNum, iterationFile, mMatrixFile, pMatrixFile string) {
	f, err := os.Open(iterationFile)
	errorCheck(err)
	iterationReader := assemble.StraightScannerIterationReader{Reader: bufio.NewReader(f)}

	remainNodes := make([]int, len(graphConfig.Nodes))
	counter := 0
	for k := range graphConfig.Nodes {
		remainNodes[counter] = k
		counter++
	}

	assemble.Assemble(M, P, remainNodes, idToNum, iterationReader)

	writeFloatMatrix(M, mMatrixFile)
	writeIntMatrix(P, pMatrixFile)
}

func revertMap(ogMap map[int]int) map[int]int {
	result := make(map[int]int, len(ogMap))
	for k, v := range ogMap {
		result[v] = k
	}
	return result
}

func reWritePmatrix(P [][]int, idToNum disassemble.IdToNum, pMatrixFile string) {
	reverter := revertMap(idToNum)
	f, err := os.Create(pMatrixFile)
	errorCheck(err)

	str := "\t"
	for i := range P {
		str += fmt.Sprintf("\t%d", reverter[i])
	}
	str += "\n"
	f.WriteString(str)

	for i := range P {
		str = fmt.Sprintf("%d\t", reverter[i])
		for j := range P[i] {
			id, ok := reverter[P[i][j]]
			if !ok {
				str += fmt.Sprintf("\t%d", P[i][j])
			} else {
				str += fmt.Sprintf("\t%d", reverter[id])
			}

		}
		str += "\n"
		f.WriteString(str)
	}
}

func main() {
	graphConfig, err := graph.Read(testFile)
	errorCheck(err)
	//fmt.Printf("%#v\n", graphConfig)

	idToNum := makeIdToNum(graphConfig, idToNumFile)

	doDisassemble(graphConfig, iterationFile, resultFile)

	M, P := microsolution.SolveOnePath(graphConfig, idToNum, 1, 2)

	doAssemble(M, P, graphConfig, idToNum, iterationFile, mMatrixFile, pMatrixFile)

	reWritePmatrix(P, idToNum, "./results/P2.txt")
}
