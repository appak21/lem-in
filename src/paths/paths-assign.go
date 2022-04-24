package paths

import (
	"container/list"
)

type Paths struct {
	Npaths, Nsteps int
	AllPaths       []**list.List
	Assignment     []int //groups of ants for each path
}

func (paths *Paths) pathLen(i int) int {
	return (*paths.AllPaths[i]).Len()
}

func (paths *Paths) calcSteps(nants int) int {
	l := len(paths.AllPaths) - 1
	shortest := paths.pathLen(0)
	longest := paths.pathLen(l)
	var sum int
	for i := 0; i < paths.Npaths; i++ {
		sum += longest - paths.pathLen(i)
	}
	numOfAnts := longest - shortest + (nants-sum)/paths.Npaths
	rem := (nants - sum) % paths.Npaths
	if rem > 0 {
		numOfAnts++
	}
	return shortest + numOfAnts - 1
}

func (paths *Paths) antsSplit(nants int) {
	paths.Assignment = make([]int, paths.Npaths)
	l := len(paths.AllPaths) - 1
	longest := paths.pathLen(l)
	var sum int
	for i := 0; i < paths.Npaths; i++ {
		sum += longest - paths.pathLen(i)
	}
	fn := float32(nants-sum) / float32(paths.Npaths)
	remSteps := (fn - float32(int(fn))) * float32(paths.Npaths)
	for i := 0; i < paths.Npaths; i++ {
		paths.Assignment[i] = longest - paths.pathLen(i) + int(fn)
		if remSteps > 0 {
			paths.Assignment[i]++
			remSteps--
		}
	}
}
