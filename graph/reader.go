package graph

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/FackOff25/disassemble_assemble/heuristics"
)

type HeuristicConf struct {
	Function string `json:"func"`
}

type GraphConfig struct {
	Heuristic HeuristicConf `json:"heuristic"`
	Nodes     []NodeConfig        `json:"graph"`
}

type NodeConfig struct {
	Id     int                    `json:"id"`
	Edges  []Edge           `json:"paths"`
	Params map[string]interface{} `json:"params"`
}

type Graph struct {
	Nodes map[int]Node
}


var graph Graph
var heuristic heuristics.Heuristic

func Read(configPath string) (Graph, error) {
	jsonFile, err := os.Open(configPath)

	if err != nil {
		return Graph{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	conf := GraphConfig{}
	json.Unmarshal(byteValue, &conf)

	graph.Nodes = make(map[int]Node)
	for _, v := range conf.Nodes {
		node := Node{
			Id: v.Id,
			Edges: make(map[int]Edge),
			Params: v.Params,
		}
		for _, edgeConf := range v.Edges {
			edge := Edge{Source: v.Id,
				Target: edgeConf.Target,
				Weight: edgeConf.Weight}
			node.Edges[edgeConf.Target] = edge
		}
		graph.Nodes[v.Id] = node
	}

	heuristic = heuristics.GetHeuristic(conf.Heuristic.Function)

	return graph, nil
}
