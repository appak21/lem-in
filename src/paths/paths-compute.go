package paths

import (
	"container/list"
	"lemin/src/algo"
	"lemin/src/antfarm"
	"sort"
)

func PathsCompute(graph *antfarm.Graph) *Paths {
	var pathsOld *Paths
	var pathsNew *Paths
	if pathsOld = PathsGetNext(graph); pathsOld == nil {
		return nil
	}
	nPaths := 1
	for nPaths < graph.Nants {
		if pathsNew = PathsGetNext(graph); pathsNew == nil {
			break
		}
		if pathsNew.Nsteps < pathsOld.Nsteps {
			pathsOld = pathsNew
		}
		nPaths++
	}
	return pathsOld
}

func PathsGetNext(graph *antfarm.Graph) *Paths {
	if !algo.Suurballe(graph) {
		return nil
	}
	return PathsFromGraph(graph)
}

func PathsFromGraph(graph *antfarm.Graph) *Paths {
	paths := new(Paths)
	paths.Npaths = graph.Exits.Len()
	paths.AllPaths = make([]**list.List, paths.Npaths)
	i := 0
	for link := graph.Exits.Front(); link != nil; link = link.Next() {
		p := unrollPath(graph, link.Value.(string))
		paths.AllPaths[i] = &p
		i++
	}
	sort.Slice(paths.AllPaths, func(i, j int) bool { return (*paths.AllPaths[i]).Len() < (*paths.AllPaths[j]).Len() })
	paths.Nsteps = paths.calcSteps(graph.Nants)
	return paths
}

func unrollPath(graph *antfarm.Graph, v string) *list.List {
	path := list.New()
	path.PushFront(graph.End)
	for v != graph.Start {
		path.PushFront(v)
		v = (*graph.Rooms[v]).Parent
	}
	return path
}
