package microsolution

import (
	"fmt"

	"github.com/FackOff25/disassemble_assemble/assemble"
	"github.com/FackOff25/disassemble_assemble/astar"
	"github.com/FackOff25/disassemble_assemble/graph"
)

func SolveOnePath(searchingGraph graph.Graph, sourceNode int, targetNode int, size int) ([][]float64, [][]int) {
	M := initMatrixM(size)
	P := initMatrixP(size)

	path, _, found := astar.Path(searchingGraph.Nodes[sourceNode], searchingGraph.Nodes[targetNode])
	if !found {
		return M, P
	}

	for i, v := range path[1:] {
		fmt.Println(i)
		for j := 0; j < i+1; j++ {
			P[path[j].GetId()-1][v.GetId()-1] = path[i].GetId()
			M[path[j].GetId()-1][v.GetId()-1] = M[path[j].GetId()-1][path[i].GetId()-1] + searchingGraph.Nodes[path[i].GetId()].Edges[v.GetId()].Weight
		}
	}

	return M, P
}

func initMatrixM(size int) [][]float64 {
	M := make([][]float64, size)
	for i := range M {
		M[i] = make([]float64, size)
		for j := range M[i] {
			if i == j {
				M[i][j] = 0
			} else {
				M[i][j] = assemble.INF
			}
		}
	}
	return M
}

func initMatrixP(size int) [][]int {
	P := make([][]int, size)
	for i := range P {
		P[i] = make([]int, size)
		for j := range P[i] {
			P[i][j] = assemble.INF
		}
	}

	return P
}
