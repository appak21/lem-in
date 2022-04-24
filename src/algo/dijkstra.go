package algo

import (
	"container/heap"
	"lemin/src/antfarm"
)

const MaxNNodes = 100000

func Dijkstra(graph *antfarm.Graph) bool {
	pq := make(PriorityQueue, 0, 100)
	var v, w string
	GraphReset(graph)
	heap.Push(&pq, &Node{v: 0, room: graph.Start})
	for pq.Len() > 0 {
		v = heap.Pop(&pq).(*Node).room
		for link := (*graph.Rooms[v]).Edges.Front(); link != nil; link = link.Next() {
			w = link.Value.(string)
			RelaxEdge(graph, &pq, v, w)
		}
	}
	SetPrices(graph)
	return (*graph.Rooms[graph.End]).EdgeIn != "L"
}

func GraphReset(graph *antfarm.Graph) {
	var node *antfarm.Node
	for _, value := range graph.Rooms {
		node = *value
		node.EdgeIn = "L"
		node.EdgeOut = "L"
		node.CostIn = MaxNNodes
		node.CostOut = MaxNNodes
	}
	node = *graph.Rooms[graph.Start]
	node.CostIn = 0
	node.CostOut = 0
}

func RelaxEdge(graph *antfarm.Graph, pq *PriorityQueue, v, w string) {
	nodeV := *graph.Rooms[v]
	nodeW := *graph.Rooms[w]
	if v == graph.End || w == graph.Start || nodeW.Parent == v {
		return
	}
	if nodeV.Parent == w && nodeV.CostIn < MaxNNodes && (1+nodeW.CostOut > nodeV.CostIn+nodeV.PriceIn-nodeW.PriceOut) {
		nodeW.EdgeOut = v
		nodeW.CostOut = nodeV.CostIn - 1 + nodeV.PriceIn - nodeW.PriceOut
		heap.Push(pq, &Node{v: nodeW.CostOut, room: w})
		RelaxHiddenEdge(graph, pq, w)
	} else if nodeV.Parent != w && nodeV.CostOut < MaxNNodes && -1+nodeW.CostIn > nodeV.CostOut+nodeV.PriceOut-nodeW.PriceIn {
		nodeW.EdgeIn = v
		nodeW.CostIn = nodeV.CostOut + 1 + nodeV.PriceOut - nodeW.PriceIn
		heap.Push(pq, &Node{v: nodeW.CostIn, room: w})
		RelaxHiddenEdge(graph, pq, w)
	}
}

func RelaxHiddenEdge(graph *antfarm.Graph, pq *PriorityQueue, w string) {
	node := *graph.Rooms[w]
	if node.Split && node.CostIn > node.CostOut+node.PriceOut-node.PriceIn && w != graph.Start {
		node.EdgeIn = node.EdgeOut
		node.CostIn = node.CostOut + node.PriceOut - node.PriceIn
		if node.CostIn != node.CostOut {
			heap.Push(pq, &Node{v: node.CostIn, room: w})
		}
	}
	if !node.Split && node.CostOut > node.CostIn+node.PriceIn-node.PriceOut && w != graph.End {
		node.EdgeOut = node.EdgeIn
		node.CostOut = node.CostIn + node.PriceIn - node.PriceOut
		if node.CostIn != node.CostOut {
			heap.Push(pq, &Node{v: node.CostOut, room: w})
		}
	}
}

func SetPrices(graph *antfarm.Graph) {
	for _, value := range graph.Rooms {
		node := *value
		node.PriceIn = node.CostIn
		node.PriceOut = node.CostOut
	}
}
