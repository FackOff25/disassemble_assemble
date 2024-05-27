package pruning

import "github.com/FackOff25/disassemble_assemble/graph"

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
		delete(edges, neighbour1)
		for neighbour2, _ := range edges {
			if neighbour2 == neighbour1 {
				continue
			}
			triangle := graph.Graph{Nodes: map[int]graph.Node{
				node.Id:    graph.Node{Id: node.Id, Edges: cloneEdgeMap(node.Edges), Params: node.Params},
				neighbour1: graph.Node{Id: subgraph.Nodes[neighbour1].Id, Edges: cloneEdgeMap(subgraph.Nodes[neighbour1].Edges), Params: subgraph.Nodes[neighbour1].Params},
				neighbour2: graph.Node{Id: subgraph.Nodes[neighbour2].Id, Edges: cloneEdgeMap(subgraph.Nodes[neighbour2].Edges), Params: subgraph.Nodes[neighbour2].Params},
			}}
			removedFromTriangle, addedToTriatriangle := pruneTriangle(triangle, node.Id)
			removed = append(removed, removedFromTriangle...)
			added = append(added, addedToTriatriangle...)
		}
	}
	return
}

func pruneTriangle(triangle graph.Graph, node int) (removed []graph.Edge, added []graph.Edge) {
	subgraph := triangle
	for neighbour1, edge1 := range subgraph.Nodes[node].Edges {
		delete(subgraph.Nodes[node].Edges, neighbour1)
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
