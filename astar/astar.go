package astar

import (
	"container/heap"
	"reflect"
)

type Pather interface {
	GetId() int
	PathNeighbors() []Pather
	PathNeighborCost(to Pather) float64
	PathEstimatedCost(to Pather) float64
}

type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

type nodeMap map[int]*node

func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p.GetId()]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p.GetId()] = n
	}
	return n
}

func Path(from, to Pather) (path []Pather, distance float64, found bool) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)
	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			// Found a path to the goal.
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}
			reverse(p)
			return p, current.cost, true
		}

		for _, neighbor := range current.pather.PathNeighbors() {
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
