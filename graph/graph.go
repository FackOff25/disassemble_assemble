package graph

import (
	"fmt"

	"github.com/FackOff25/disassemble_assemble/astar"
)

type Edge struct {
	Source int
	Target int
	Weight float64 `json:"weight"`
}

type Node struct {
	Id     int                    `json:"id"`
	Edges  map[int]Edge           `json:"paths"`
	Params map[string]interface{} `json:"params"`
}

func (n Node) GetId() int {
	return n.Id
}

func (n Node) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for k, _ := range n.Edges {
		neighbors = append(neighbors, graph.Nodes[k])
	}
	return neighbors
}

func (n Node) PathNeighborCost(to astar.Pather) float64 {
	for k, v := range n.Edges {
		if k == to.(Node).Id {
			return v.Weight
		}
	}
	return 0
}

func (n Node) PathEstimatedCost(to astar.Pather) float64 {
	res, _ := heuristic(n.Params, to.(Node).Params)
	return res
}

// unweighted graph comparison
func (e Edge) IsEqual(e2 Edge) bool {
	if e.Source == e2.Source && e.Target == e2.Target && e.Weight == e.Weight {
		return true
	}

	if e.Source == e2.Target && e.Target == e2.Source && e.Weight == e.Weight {
		return true
	}

	return false
}

// unweighted graph comparison
func EdgeCompare(e1 Edge, e2 Edge) bool {
	if e1.Source < e1.Target {
		if e2.Source < e2.Target {
			if e1.Source == e1.Source {
				return e1.Weight < e2.Weight
			} else {
				return e1.Source < e1.Source
			}
		} else {
			if e1.Source == e1.Target {
				return e1.Weight < e2.Weight
			} else {
				return e1.Source < e1.Target
			}
		}
	} else {
		if e2.Source < e2.Target {
			if e1.Target == e1.Source {
				return e1.Weight < e2.Weight
			} else {
				return e1.Target < e1.Source
			}
		} else {
			if e1.Target == e1.Target {
				return e1.Weight < e2.Weight
			} else {
				return e1.Target < e1.Target
			}
		}
	}
}

func IsEqualEdgeSlices(a, b []Edge) bool {
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

func IsEqualEdgeMaps(a, b map[int]Edge) bool {
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

func (n Node) IsEqual(n2 Node) bool {
	if n.Id != n2.Id {
		return false
	}
	if !IsEqualEdgeMaps(n.Edges, n2.Edges) {
		return false
	}
	return true
}

func IsEqualNodeMaps(a, b map[int]Node) bool {
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

func IsEqualGraphs(a, b Graph) bool {
	return IsEqualNodeMaps(a.Nodes, b.Nodes)
}

func (g Graph) RemoveNodes(nodes []Node) {
	for _, node := range nodes {
		for k := range node.Edges {
			delete(g.Nodes[k].Edges, node.Id)
		}
		delete(g.Nodes, node.Id)
	}
}

func (g Graph) RemoveEdges(edges []Edge) {
	for _, edge := range edges {
		node, ok := g.Nodes[edge.Source]
		if ok {
			delete(node.Edges, edge.Target)
		}
		node, ok = g.Nodes[edge.Target]
		if ok {
			delete(node.Edges, edge.Source)
		}
	}
}

func (g Graph) AddEdges(edges []Edge) {
	for _, edge := range edges {
		_, ok := g.Nodes[edge.Source]
		if !ok {
			panic(fmt.Sprintf("No node %d for edge %v", edge.Source, edge))
		}
		_, ok = g.Nodes[edge.Target]
		if !ok {
			panic(fmt.Sprintf("No node %d for edge %v", edge.Target, edge))
		}
		g.Nodes[edge.Source].Edges[edge.Target] = edge
		g.Nodes[edge.Target].Edges[edge.Source] = edge
	}
}
