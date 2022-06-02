package antfarm

import (
	"container/list"
	"fmt"
	"lemin/config"
	"math"
	"strconv"
)

const (
	AntsField = iota
	RoomsField
	LinksField
)

func (graph *Graph) ParseAnts(line string) error {
	n, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf(config.ErrAnts)
	}
	if n <= 0 || n > math.MaxInt32 {
		return fmt.Errorf(config.ErrAnts)
	}
	graph.Nants = n
	graph.data.field = RoomsField
	return nil
}

func (graph *Graph) ParseRooms(line string) error {
	if isStart(line) && !graph.data.startFound {
		graph.data.startFound = true
	} else if isEnd(line) && !graph.data.endFound {
		graph.data.endFound = true
	} else {
		room, err := getRoom(line, graph)
		if err != nil {
			return err
		}
		if graph.Rooms[room] != nil {
			return fmt.Errorf("the %v room is duplicated", room)
		}
		if room != "" {
			node := &Node{Parent: "L", Edges: make(map[string]byte)}
			if graph.data.startFound && graph.Start == "" {
				graph.Start = room
			} else if graph.data.endFound && graph.End == "" {
				graph.End = room
			}
			graph.Rooms[room] = &node
			graph.Exits = list.New()
		} else if graph.Start != "" && graph.End != "" {
			graph.data.field = LinksField
			graph.data.coordinates = nil
			return graph.ParseData(line)
		} else {
			return fmt.Errorf(config.ErrNoStart + " or " + config.ErrNoEnd)
		}
	}
	return nil
}

func (graph *Graph) ParseLinks(line string) error {
	room1, room2 := getLink(line)
	if room1 == "" && room2 == "" {
		return fmt.Errorf("invalid link")
	}
	if graph.Rooms[room1] == nil || graph.Rooms[room2] == nil {
		return fmt.Errorf("the link contains an unknown room: %v", line)
	}
	if room1 == room2 {
		return fmt.Errorf("the %v room is linked to itself: %v", room1, line)
	}
	node1 := *graph.Rooms[room1]
	node1.Edges[room2] = 1
	node2 := *graph.Rooms[room2]
	node2.Edges[room1] = 1
	return nil
}

func (graph *Graph) ParseData(line string) error {
	if isComment(line) || line == "" {
		return nil
	}
	switch graph.data.field {
	case LinksField:
		return graph.ParseLinks(line)
	case RoomsField:
		return graph.ParseRooms(line)
	case AntsField:
		return graph.ParseAnts(line)
	default:
		return fmt.Errorf("something went wrong while parsing")
	}
}
