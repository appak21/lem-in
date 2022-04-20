package paths

import (
	"container/list"
	"fmt"
)

func Lemin(paths *Paths, nants int) {
	if paths.pathLen(0) == 1 {
		room := (*paths.Arr[0]).Front().Value
		for i := 1; i <= nants; i++ {
			fmt.Printf("L%d-%v ", i, room)
		}
		fmt.Println()
		return
	}
	var ants []*list.Element
	var ants_active int
	var newLine bool
	ants = make([]*list.Element, nants)
	for ants_active < nants {
		for i := 0; i < paths.Npaths && ants_active < nants; i++ {
			if (*paths.Assignments)[i] > 0 {
				(*paths.Assignments)[i] -= 1
				ants[ants_active] = (*paths.Arr[i]).Front()
				ants_active++
			}
		}
		stepOnce(ants, nants, newLine)
		newLine = true
	}
	for stepOnce(ants, nants, newLine) > 0 {
		continue
	}
	fmt.Println()
}

func stepOnce(ants []*list.Element, nants int, newLine bool) int {
	antsMoved := 0
	for i := 0; i < nants; i++ {
		if ants[i] != nil {
			if antsMoved == 0 && newLine {
				fmt.Println()
			}
			fmt.Printf("L%d-%s ", i+1, ants[i].Value)
			ants[i] = ants[i].Next()
			antsMoved++
		}
	}
	return antsMoved
}

/*
func PathsPrint(paths *Paths, graph *antfarm.Graph) {
	var link *list.Element
	i := 0
	for i < paths.Npaths {
		link = (*paths.Arr[i]).Front()
		fmt.Print("paths: ", graph.Start, " -> ")
		for link != nil {
			fmt.Print(link.Value.(string))
			link = link.Next()
			if link != nil {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
		i++
	}
}*/
