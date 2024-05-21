package graph

import "CourseProjectA/astar"

type Edge struct {
	Target int     `json:"target"`
	Weight float64 `json:"weight"`
}

type Node struct {
	Id     int                    `json:"id"`
	Edges  []Edge                 `json:"paths"`
	Params map[string]interface{} `json:"params"`
}

func (n Node) GetId() int {
	return n.Id
}

func (n Node) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, v := range n.Edges {
		neighbors = append(neighbors, graph.Nodes[v.Target])
	}
	return neighbors
}

func (n Node) PathNeighborCost(to astar.Pather) float64 {
	for _, v := range n.Edges {
		if v.Target == to.(Node).Id {
			return v.Weight
		}
	}
	return 0
}

func (n Node) PathEstimatedCost(to astar.Pather) float64 {
	res, _ := heuristic(n.Params, to.(Node).Params)
	return res
}
