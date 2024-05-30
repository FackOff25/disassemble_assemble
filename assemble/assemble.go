package assemble

import (
	"github.com/FackOff25/disassemble_assemble/disassemble"
	"github.com/FackOff25/disassemble_assemble/graph"
)

const (
	INF = 99999999
)

type IterationChanges struct {
	AddedNodes   []graph.Node
	RemovedEdges []disassemble.Edge
}

type IteractionReader interface {
	ReadNextIteration() IterationChanges
}

func Assemble(M [][]float64, P [][]int, nodes []int, reader IteractionReader) {
	for len(M) < len(nodes) {
		iterationChanges := reader.ReadNextIteration()
		for _, v := range iterationChanges.AddedNodes {
			for k, e := range v.Edges {
				M[v.Id][k] = e.Weight
				M[k][v.Id] = e.Weight
				P[v.Id][k] = v.Id
				P[k][v.Id] = k
			}
			for _, node := range nodes {
				if M[v.Id][node] == INF {
					for k, _ := range v.Edges {
						if M[v.Id][node] > M[k][node]+M[v.Id][k] {
							M[v.Id][k] = M[k][node] + M[v.Id][k]
							M[k][v.Id] = M[k][node] + M[v.Id][k]
							P[v.Id][node] = P[k][node]
							P[node][v.Id] = node
						}
					}
				}
			}
		}
		for _, edge := range iterationChanges.RemovedEdges {
			previousnessLine := P[edge.TheEdge.Target]
			for _, node := range nodes {
				if previousnessLine[node] == edge.TheEdge.Source {
					previousnessLine[node] = edge.LinkedNode
				}
			}
			for _, v := range iterationChanges.AddedNodes {
				if previousnessLine[v.Id] == edge.TheEdge.Source {
					previousnessLine[v.Id] = edge.LinkedNode
				}
			}
		}
	}
}
