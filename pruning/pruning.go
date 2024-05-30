package pruning

import (
	"github.com/FackOff25/disassemble_assemble/graph"
)

func cloneEdgeMap(ogMap map[int]graph.Edge) map[int]graph.Edge {
	result := make(map[int]graph.Edge)
	for k, v := range ogMap {
		result[k] = v
	}
	return result
}

func GetNeighbourSubgraph(wholeGraph graph.Graph, node graph.Node) graph.Graph {
	subgraph := graph.Graph{Nodes: map[int]graph.Node{node.Id: node}}
	for id, _ := range node.Edges {
		subgraph.Nodes[id] = wholeGraph.Nodes[id]
	}
	return subgraph
}

func RemoveNode(subgraph graph.Graph, node graph.Node) (removed []graph.Edge, added []graph.Edge) {
	edges := cloneEdgeMap(node.Edges)
	for neighbour1, _ := range edges {
		for neighbour2, _ := range edges {
			if neighbour2 == neighbour1 || neighbour2 < neighbour1 {
				continue
			}
			removedFromTriangle, addedToTriatriangle := pruneTriangle(subgraph, node.Id, neighbour1, neighbour2)
			removed = append(removed, removedFromTriangle...)
			added = append(added, addedToTriatriangle...)
		}
	}
	return
}

func pruneTriangle(subgraph graph.Graph, mainVertex, sideVertex1, sideVertex2 int) (removed []graph.Edge, added []graph.Edge) {
	removed = make([]graph.Edge, 0)
	added = make([]graph.Edge, 0)

	edge, ok := subgraph.Nodes[sideVertex1].Edges[sideVertex2]
	edge1 := subgraph.Nodes[sideVertex1].Edges[mainVertex]
	edge2 := subgraph.Nodes[mainVertex].Edges[sideVertex2]
	sumWeight := edge1.Weight + edge2.Weight
	if ok {
		if edge.Weight > sumWeight {
			removed = append(removed, edge)
			added = append(added, graph.Edge{Source: sideVertex1, Target: sideVertex2, Weight: sumWeight})
		}
	} else {
		added = append(added, graph.Edge{Source: sideVertex1, Target: sideVertex2, Weight: sumWeight})
	}
	return
}
