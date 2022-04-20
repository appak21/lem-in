package antfarm

import (
	"fmt"
	"lemin/config"
	"strconv"
	"strings"
)

func isComment(line string) bool {
	return (strings.HasPrefix(line, "#") && !isStart(line) && !isEnd(line))
}

func isStart(room string) bool {
	return room == "##start"
}

func isEnd(room string) bool {
	return room == "##end"
}

func isRoom(room string) bool {
	return !strings.HasPrefix(room, "L") //&& len(room) > 0
}

func isCoord(x, y string, graph *Graph) bool {
	xn, err1 := strconv.Atoi(x)
	yn, err2 := strconv.Atoi(y)
	if err1 != nil || err2 != nil {
		return false
	}
	if _, ok := graph.data.coordinates[[2]int{xn, yn}]; ok {
		return false
	}
	graph.data.coordinates = make(map[[2]int]bool)
	graph.data.coordinates[[2]int{xn, yn}] = true
	return true
}

func getRoom(room string, graph *Graph) (string, error) {
	s := strings.Split(room, " ")
	if len(s) != 3 {
		return "", nil //nil, because it should be able to move on if Link's field starts
	}
	if !isRoom(s[0]) {
		return "", fmt.Errorf(config.ErrRoomName)
	}
	if !isCoord(s[1], s[2], graph) {
		return "", fmt.Errorf(config.ErrCoord+": duplicated %v", room)
	}
	return s[0], nil
}

func getLink(link string) (string, string) {
	s := strings.Split(link, "-")
	if len(s) != 2 {
		return "", ""
	}
	return s[0], s[1]
}
