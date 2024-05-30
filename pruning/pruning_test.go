package pruning

import (
	"sort"
	"testing"

	"github.com/FackOff25/disassemble_assemble/graph"
)

type pruneTestCase struct {
	Subgraph graph.Graph
	Node     int
	Removed  []graph.Edge
	Added    []graph.Edge
}

func sortEdgeSlice(slice []graph.Edge) {
	sort.Slice(slice, func(i, j int) bool {
		e1 := slice[i]
		e2 := slice[j]
		e1Min := min(e1.Source, e1.Target)
		e1Max := max(e1.Source, e1.Target)

		e2Min := min(e2.Source, e2.Target)
		e2Max := max(e2.Source, e2.Target)

		if e1Min == e2Min {
			if e1Max == e2Max {
				return e1.Weight < e2.Weight
			}
			return e1Max < e2Max
		}
		return e1Min < e2Min
	})
}

type trianglePruneTestCase struct {
	Subgraph    graph.Graph
	MainVertex  int
	SideVertex1 int
	SideVertex2 int
	Removed     []graph.Edge
	Added       []graph.Edge
}

func TestTrianglePrune(t *testing.T) {
	cases := []trianglePruneTestCase{
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
		}}, MainVertex: 1, SideVertex1: 2, SideVertex2: 3, Removed: []graph.Edge{}, Added: []graph.Edge{}},
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 3, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
		}}, MainVertex: 1, SideVertex1: 2, SideVertex2: 3, Removed: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 6}}, Added: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 2}}},
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
		}}, MainVertex: 1, SideVertex1: 2, SideVertex2: 3, Removed: []graph.Edge{}, Added: []graph.Edge{graph.Edge{Source: 2, Target: 3, Weight: 2}}},
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 4}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 3, Weight: 4}, 2: graph.Edge{Source: 3, Target: 2, Weight: 1}}, Params: map[string]interface{}{}},
		}}, MainVertex: 1, SideVertex1: 2, SideVertex2: 3, Removed: []graph.Edge{}, Added: []graph.Edge{}},
	}
	for i, testCase := range cases {
		sortEdgeSlice(testCase.Added)
		sortEdgeSlice(testCase.Removed)
		resultRemoved, resultAdded := pruneTriangle(testCase.Subgraph, testCase.MainVertex, testCase.SideVertex1, testCase.SideVertex2)
		sortEdgeSlice(resultRemoved)
		sortEdgeSlice(resultAdded)
		if !graph.IsEqualEdgeSlices(resultRemoved, testCase.Removed) {
			t.Errorf("Wrong read in %d case. Expected removed: %v, got removed: %v", i, testCase.Removed, resultRemoved)
		}
		if !graph.IsEqualEdgeSlices(resultAdded, testCase.Added) {
			t.Errorf("Wrong read in %d case. Expected added: %v, got added: %v", i, testCase.Added, resultAdded)
		}
	}
}

func TestPrune(t *testing.T) {
	cases := []pruneTestCase{
		{Subgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}, 4: graph.Edge{Source: 1, Target: 4, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 1}, 4: graph.Edge{Source: 4, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
			4: graph.Node{Id: 4, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 4, Weight: 1}, 3: graph.Edge{Source: 4, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
		}}, Node: 1, Removed: []graph.Edge{graph.Edge{Source: 4, Target: 3, Weight: 6}}, Added: []graph.Edge{graph.Edge{Source: 4, Target: 2, Weight: 2}, graph.Edge{Source: 4, Target: 3, Weight: 2}}},
	}
	for i, testCase := range cases {
		sortEdgeSlice(testCase.Added)
		sortEdgeSlice(testCase.Removed)
		resultRemoved, resultAdded := RemoveNode(testCase.Subgraph, testCase.Subgraph.Nodes[testCase.Node])
		sortEdgeSlice(resultRemoved)
		sortEdgeSlice(resultAdded)
		if !graph.IsEqualEdgeSlices(resultRemoved, testCase.Removed) {
			t.Errorf("Wrong read in %d case. Expected removed: %v, got removed: %v", i, testCase.Removed, resultRemoved)
		}
		if !graph.IsEqualEdgeSlices(resultAdded, testCase.Added) {
			t.Errorf("Wrong read in %d case. Expected added: %v, got added: %v", i, testCase.Added, resultAdded)
		}
	}
}

type getNeighbourSubgraphTestCase struct {
	WholeGraph       graph.Graph
	Node             int
	ExpectedSubgraph graph.Graph
}

func TestGetSubgraph(t *testing.T) {
	cases := []getNeighbourSubgraphTestCase{
		{WholeGraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}, 4: graph.Edge{Source: 2, Target: 4, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 1}, 4: graph.Edge{Source: 4, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
			4: graph.Node{Id: 4, Edges: map[int]graph.Edge{2: graph.Edge{Source: 2, Target: 4, Weight: 1}, 3: graph.Edge{Source: 4, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
		}}, Node: 1, ExpectedSubgraph: graph.Graph{Nodes: map[int]graph.Node{
			1: graph.Node{Id: 1, Edges: map[int]graph.Edge{2: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 1, Target: 3, Weight: 1}}, Params: map[string]interface{}{}},
			2: graph.Node{Id: 2, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 3: graph.Edge{Source: 2, Target: 3, Weight: 1}, 4: graph.Edge{Source: 2, Target: 4, Weight: 1}}, Params: map[string]interface{}{}},
			3: graph.Node{Id: 3, Edges: map[int]graph.Edge{1: graph.Edge{Source: 1, Target: 2, Weight: 1}, 2: graph.Edge{Source: 2, Target: 3, Weight: 1}, 4: graph.Edge{Source: 4, Target: 3, Weight: 6}}, Params: map[string]interface{}{}},
		}},
		},
	}
	for i, testCase := range cases {
		resultSubgraph := GetNeighbourSubgraph(testCase.WholeGraph, testCase.WholeGraph.Nodes[testCase.Node])
		if !graph.IsEqualGraphs(resultSubgraph, testCase.ExpectedSubgraph) {
			t.Errorf("Wrong read in %d case. Expected removed: %v, got removed: %v", i, testCase.ExpectedSubgraph, resultSubgraph)
		}
	}
}
