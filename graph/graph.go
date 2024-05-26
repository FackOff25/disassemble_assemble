package graph

import "github.com/FackOff25/disassemble_assemble/astar"

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
