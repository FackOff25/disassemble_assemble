package graph

import (
	"CourseProjectA/heuristics"
	"encoding/json"
	"io/ioutil"
	"os"
)

type HeuristicConf struct {
	Function string `json:"func"`
}

type GraphConfig struct {
	Heuristic HeuristicConf `json:"heuristic"`
	Nodes     []Node        `json:"graph"`
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
		graph.Nodes[v.Id] = v
	}

	heuristic = heuristics.GetHeuristic(conf.Heuristic.Function)

	return graph, nil
}
