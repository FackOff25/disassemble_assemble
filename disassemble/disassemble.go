package disassemble

import (
	"encoding/json"

	"github.com/FackOff25/disassemble_assemble/graph"
	"github.com/FackOff25/disassemble_assemble/pruning"
)

type VertexChoseStrategy interface {
	ChoseVertexes(_graph graph.Graph, numberToChose int) []graph.Node
}

type PruningEndStrategy interface {
	// True if time to stop
	CheckPruningEnd(_graph graph.Graph) bool
}

type IterationWriter interface {
	Write(IterationChanges)
}

type Edge struct {
	LinkedNode int        `json:"LinkedNode"`
	TheEdge    graph.Edge `json:"Edge"`
}

type IterationChanges struct {
	RemovedNodes []graph.Node `json:"removedNodes"`
	RemovedEdges []Edge       `json:"removedEdges"`
	AddedEdges   []Edge       `json:"addedEdges"`
}

type IdToNum map[int]int

func MakeIdToNum(originalGraph graph.Graph) IdToNum {
	result := make(IdToNum)
	counter := 0
	for k := range originalGraph.Nodes {
		result[k] = counter
		counter++
	}
	return result
}

func (ic IterationChanges) ToString() string {
	byteStr, _ := json.Marshal(ic)
	return string(byteStr)
}

func transformEdgeSlice(slice []graph.Edge, node int) []Edge {
	result := make([]Edge, len(slice))
	for i, v := range slice {
		result[i] = Edge{TheEdge: v, LinkedNode: node}
	}
	return result
}

func Disassemble(originalGraph graph.Graph, nodeChoser VertexChoseStrategy, endPruneStrategy PruningEndStrategy, iterationWriter IterationWriter) {
	iteration := 0
	for !endPruneStrategy.CheckPruningEnd(originalGraph) {
		iteration++
		removingNodes := nodeChoser.ChoseVertexes(originalGraph, 1)
		iterationRemovedEdges := make([]Edge, 0)
		removingEdges := make([]graph.Edge, 0)
		iterationAddedEdges := make([]Edge, 0)
		addingEdges := make([]graph.Edge, 0)
		removingNodesEdges := make([]Edge, 0)

		for _, removingNode := range removingNodes {
			pruningSubgraph := pruning.GetNeighbourSubgraph(originalGraph, removingNode)
			removedEdges, addedEdges := pruning.RemoveNode(pruningSubgraph, removingNode)
			removingEdges = append(removingEdges, removedEdges...)
			iterationRemovedEdges = append(iterationRemovedEdges, transformEdgeSlice(removedEdges, removingNode.Id)...)
			addingEdges = append(addingEdges, addedEdges...)
			iterationAddedEdges = append(iterationAddedEdges, transformEdgeSlice(addedEdges, removingNode.Id)...)
			for _, v := range removingNode.Edges {
				removingNodesEdges = append(removingNodesEdges, Edge{TheEdge: v, LinkedNode: removingNode.Id})
			}
		}

		originalGraph.RemoveNodes(removingNodes)
		originalGraph.RemoveEdges(removingEdges)
		originalGraph.AddEdges(addingEdges)

		iterationWriter.Write(IterationChanges{RemovedNodes: removingNodes, RemovedEdges: iterationRemovedEdges, AddedEdges: iterationAddedEdges})
		//time.Sleep(2 + time.Second)
	}
}
