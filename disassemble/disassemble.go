package disassemble

import (
	"encoding/json"
	"fmt"
	"os"

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

type IterationChanges struct {
	RemovedNodes []int        `json:"rn"`
	RemovedEdges []graph.Edge `json:"re"`
	AddedEdges   []graph.Edge `json:"ae"`
}

func (ic IterationChanges) ToString() string {
	byteStr, _ := json.Marshal(ic)
	return string(byteStr)
}

func Disassemble(originalGraph graph.Graph, nodeChoser VertexChoseStrategy, endPruneStrategy PruningEndStrategy, iterationWriter IterationWriter) {
	iteration := 0
	for !endPruneStrategy.CheckPruningEnd(originalGraph) {
		iteration++
		removingNodes := nodeChoser.ChoseVertexes(originalGraph, 1)
		iterationRemovedEdges := make([]graph.Edge, 0)
		iterationAddedEdges := make([]graph.Edge, 0)

		r, err := os.Create("./results/iteration" + fmt.Sprint(iteration) + ".json")
		if err != nil {
			fmt.Errorf("Error: %s", err)
			return
		}
		byteStr, _ := json.Marshal(originalGraph)
		r.Write(byteStr)
		r.Close()

		for _, removingNode := range removingNodes {
			pruningSubgraph := pruning.GetNeighbourSubgraph(originalGraph, removingNode)
			removedEdges, addedEdges := pruning.RemoveNode(pruningSubgraph, removingNode)
			iterationRemovedEdges = append(iterationRemovedEdges, removedEdges...)
			iterationAddedEdges = append(iterationAddedEdges, addedEdges...)
		}

		originalGraph.RemoveNodes(removingNodes)
		originalGraph.RemoveEdges(iterationRemovedEdges)
		originalGraph.AddEdges(iterationAddedEdges)

		removingNodesInts := make([]int, 0)
		for _, v := range removingNodes {
			removingNodesInts = append(removingNodesInts, v.Id)
		}

		iterationWriter.Write(IterationChanges{RemovedNodes: removingNodesInts, RemovedEdges: iterationRemovedEdges, AddedEdges: iterationAddedEdges})
		//time.Sleep(2 + time.Second)
	}
}
