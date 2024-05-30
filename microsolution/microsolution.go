package microsolution

import (
	"github.com/FackOff25/disassemble_assemble/assemble"
	"github.com/FackOff25/disassemble_assemble/astar"
	"github.com/FackOff25/disassemble_assemble/disassemble"
	"github.com/FackOff25/disassemble_assemble/graph"
)

func SolveOnePath(searchingGraph graph.Graph, idToNum disassemble.IdToNum, sourceNode int, targetNode int) ([][]float64, [][]int) {
	M := initMatrixM(len(idToNum))
	P := initMatrixP(len(idToNum))

	path, _, found := astar.Path(searchingGraph.Nodes[sourceNode], searchingGraph.Nodes[targetNode])
	if !found {
		return M, P
	}

	for i, v := range path[1:] {
		for j := 0; j < i+1; j++ {
			pathJNum := idToNum[path[j].GetId()]
			pathINum := idToNum[path[i].GetId()]
			vNum := idToNum[v.GetId()]

			P[pathJNum][vNum] = pathINum
			P[vNum][pathJNum] = idToNum[path[j+1].GetId()]
			M[pathJNum][vNum] = M[pathJNum][pathINum] + searchingGraph.Nodes[path[i].GetId()].Edges[v.GetId()].Weight
			M[vNum][pathJNum] = M[pathJNum][vNum]
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
