package pruning

import "github.com/FackOff25/disassemble_assemble/graph"

func RemoveNode(subgraph graph.Graph, node graph.Node) (removed graph.Graph, added []graph.Edge) {
	return
}

func pruneTriangle(subgraph graph.Graph, node int) (removed []graph.Edge, added []graph.Edge) {
	for neighbour1, edge1 := range subgraph.Nodes[node].Edges {
		for neighbour2, edge2 := range subgraph.Nodes[node].Edges {
			if neighbour2 == neighbour1 {
				continue
			}
			edge, ok := subgraph.Nodes[neighbour1].Edges[neighbour2]
			if ok {
				if edge.Weight > edge1.Weight+edge2.Weight {
					removed = append(removed, edge)
					added = append(added, graph.Edge{Source: neighbour1, Target: neighbour2, Weight: edge1.Weight + edge2.Weight})
				}
			} else {
				added = append(added, graph.Edge{Source: neighbour1, Target: neighbour2, Weight: edge1.Weight + edge2.Weight})
			}
		}
	}
	return
}
