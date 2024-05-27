package disassemble

import (
	"slices"

	"github.com/FackOff25/disassemble_assemble/graph"
)

type RandomChoser struct {
	ExcludeNodes []int
}

func (rc RandomChoser) ChoseVertexes(_graph graph.Graph, numberToChose int) []graph.Node {
	counter := 0
	chosen := make([]graph.Node, numberToChose)
	for k, v := range _graph.Nodes {
		if slices.Contains(rc.ExcludeNodes, k) {
			continue
		}
		chosen[counter] = v
		counter++
	}
	return chosen
}
