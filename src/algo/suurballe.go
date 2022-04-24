package algo

import (
	"lemin/src/antfarm"
)

func Suurballe(graph *antfarm.Graph) bool {
	if !Dijkstra(graph) {
		return false
	}
	CachePath(graph)
	return true
}

func CachePath(graph *antfarm.Graph) {
	var unsplit bool
	w := graph.End
	v := (*graph.Rooms[w]).EdgeIn
	graph.Exits.PushBack(v)
	for w != graph.Start {
		if (*graph.Rooms[v]).Parent == w {
			if unsplit {
				unsplitNode(graph, w)
			}
			unsplit = true
			simultAssign(&w, &v, v, (*graph.Rooms[v]).EdgeIn)
		} else {
			(*graph.Rooms[w]).Parent = v
			splitNode(graph, w)
			unsplit = false
			simultAssign(&w, &v, v, (*graph.Rooms[v]).EdgeOut)
		}
	}
}

func unsplitNode(graph *antfarm.Graph, v string) {
	(*graph.Rooms[v]).Split = false
	(*graph.Rooms[v]).Parent = "L"
}

func splitNode(graph *antfarm.Graph, v string) {
	if v != graph.Start && v != graph.End {
		(*graph.Rooms[v]).Split = true
	}
}

func simultAssign(a, b *string, c, d string) {
	*a = c
	*b = d
}
