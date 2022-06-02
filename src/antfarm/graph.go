package antfarm

import "container/list"

type Graph struct {
	Rooms      map[string]**Node
	Exits      *list.List
	Start, End string
	Nants      int
	data       parseInfo
}

type parseInfo struct {
	coordinates          map[[2]int]bool
	field                byte
	startFound, endFound bool
}

type Node struct {
	Edges             map[string]byte
	Parent            string
	EdgeIn, EdgeOut   string
	PriceIn, PriceOut int
	CostIn, CostOut   int
	Split             bool
}

func NewGraph() *Graph {
	return &Graph{Rooms: make(map[string]**Node)}
}
