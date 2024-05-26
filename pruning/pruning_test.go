package pruning

import (
	"testing"

	"github.com/FackOff25/disassemble_assemble/graph"
)

func isEqualEdgeSlices(a, b []graph.Edge) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !v.IsEqual(b[i]) {
			return false
		}
	}
	return true
}

type trianglePruneTestCase struct {
	Subgraph graph.Graph
	Node     int
	Removed  []graph.Edge
	Added    []graph.Edge
}

func TestTrianglePrune(t *testing.T) {
	cases := []trianglePruneTestCase{
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
		}}, Node: 1, Removed: []graph.Edge{}, Added: []graph.Edge{}},
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
		}}, Node: 1, Removed: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 6}, graph.Edge{Source: 3, Target: 2, Weight: 6}}, Added: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 2}, graph.Edge{Source: 3, Target: 2, Weight: 2}}},
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}}, Params: map[string]interface{}{}},
		}}, Node: 1, Removed: []graph.Edge{}, Added: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 2}, graph.Edge{Source: 3, Target: 2, Weight: 2}}},
	}
	for i, testCase := range cases {
		resultRemoved, resultAdded := pruneTriangle(testCase.Subgraph, testCase.Node)
		if !isEqualEdgeSlices(resultRemoved, testCase.Removed) {
			t.Errorf("Wrong read in %d case. Expected removed: %v, got removed: %v", i, testCase.Removed, resultRemoved)
		}
		if !isEqualEdgeSlices(resultAdded, testCase.Added) {
			t.Errorf("Wrong read in %d case. Expected added: %v, got added: %v", i, testCase.Added, resultAdded)
		}
	}
}
