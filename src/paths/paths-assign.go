package paths

import (
	"container/list"
	"math"
	"sort"
)

type Paths struct {
	Npaths, Nsteps int
	Arr            []**list.List //all paths
	Assignments    *[]int        //the number of ants that should take each path
}

func (paths *Paths) pathLen(i int) int {
	return (*paths.Arr[i]).Len()
}

func (paths *Paths) pathsAssign(nants int) {
	var nStepsOld, nStepsNew int
	var assignmentsNew *[]int
	pA := make([]int, nants)
	assignmentsNew = &pA
	nStepsOld = math.MaxInt
	sort.Slice(paths.Arr, func(i, j int) bool { return (*paths.Arr[i]).Len() < (*paths.Arr[j]).Len() })
	i := 1
	for i <= paths.Npaths && nStepsOld > paths.pathLen(i-1) {
		nStepsNew = paths.pathsAssignOnce(nants, i, assignmentsNew)
		paths.pathsReassign(assignmentsNew)
		nStepsOld = nStepsNew
		i++
	}
	paths.Nsteps = nStepsOld
}

func (paths *Paths) pathsAssignOnce(nants, nused int, assignmentsNew *[]int) int {
	var nStepsNew, i int
	for i+1 < nused {
		nants -= paths.pathLen(nused-1) - paths.pathLen(i)
		if paths.pathLen(nused-1) > paths.pathLen(i) {
			(*assignmentsNew)[i] = paths.pathLen(nused-1) - paths.pathLen(i)
		} else {
			(*assignmentsNew)[i] = 0
		}
		i++
	}
	i = 0
	for i < nused {
		(*assignmentsNew)[i] += nants / nused
		if i < nants%nused {
			(*assignmentsNew)[i] += 1
		}
		i++
	}
	nStepsNew = paths.pathLen(0) + (*assignmentsNew)[0] - 1
	return nStepsNew
}

func (paths *Paths) pathsReassign(assignmentsNew *[]int) {
	i := 0
	for i < paths.Npaths {
		(*paths.Assignments)[i] = (*assignmentsNew)[i]
		i++
	}
}
