package assemble

import (
	"github.com/FackOff25/disassemble_assemble/disassemble"
	"github.com/FackOff25/disassemble_assemble/graph"
)

const (
	INF = 99
)

type IterationChanges struct {
	AddedNodes   []graph.Node
	RemovedEdges []disassemble.Edge
}

type IteractionReader interface {
	ReadNextIteration() IterationChanges
}

func Assemble(M [][]float64, P [][]int, nodes []int, idToNum disassemble.IdToNum, reader IteractionReader) {
	for len(M) > len(nodes) {
		iterationChanges := reader.ReadNextIteration()
		for _, v := range iterationChanges.AddedNodes {
			vNum := idToNum[v.Id]
			for k, e := range v.Edges {
				kNum := idToNum[k]

				M[vNum][kNum] = e.Weight
				M[kNum][vNum] = e.Weight
				P[vNum][kNum] = vNum
				P[kNum][vNum] = kNum
			}
			for _, node := range nodes {
				nodeNum := idToNum[node]
				if M[vNum][nodeNum] == INF {
					for k := range v.Edges {
						kNum := idToNum[k]
						sumWeight := M[vNum][kNum] + M[kNum][nodeNum]
						if M[vNum][nodeNum] > sumWeight {
							M[vNum][nodeNum] = sumWeight
							M[nodeNum][vNum] = sumWeight
							P[vNum][nodeNum] = P[kNum][nodeNum]
							P[nodeNum][vNum] = kNum
						}
					}
				}
			}
		}
		for _, edge := range iterationChanges.RemovedEdges {
			sourceNum := idToNum[edge.TheEdge.Source]
			targetNum := idToNum[edge.TheEdge.Target]
			linkedNum := idToNum[edge.LinkedNode]

			for _, node := range nodes {
				nodeNum := idToNum[node]
				if P[nodeNum][targetNum] == sourceNum {
					P[nodeNum][targetNum] = linkedNum
				}
			}
			for _, v := range iterationChanges.AddedNodes {
				vNum := idToNum[v.Id]
				if P[vNum][targetNum] == sourceNum {
					P[vNum][targetNum] = linkedNum
				}
			}

			for _, node := range nodes {
				nodeNum := idToNum[node]
				if P[nodeNum][sourceNum] == targetNum {
					P[nodeNum][sourceNum] = linkedNum
				}
			}
			for _, v := range iterationChanges.AddedNodes {
				vNum := idToNum[v.Id]
				if P[vNum][sourceNum] == targetNum {
					P[vNum][sourceNum] = linkedNum
				}
			}
		}
		addedNodesInt := make([]int, len(iterationChanges.AddedNodes))
		for i, v := range iterationChanges.AddedNodes {
			addedNodesInt[i] = v.Id
		}
		nodes = append(nodes, addedNodesInt...)
	}
}
