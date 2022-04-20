package main

import (
	"bufio"
	"fmt"
	"lemin/config"
	"lemin/src/antfarm"
	"lemin/src/paths"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadData(filename string) (*antfarm.Graph, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf(config.ErrFileIssue)
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf(config.ErrFileIssue)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := antfarm.NewGraph()
	for scanner.Scan() {
		err := graph.ParseData(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf(config.ErrData + err.Error())
		}
	}
	return graph, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(config.ErrArgs)
		os.Exit(1)
	}
	graph, err := ReadData(args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	allPaths := paths.PathsCompute(graph)
	if allPaths == nil {
		fmt.Print(config.ErrNoPaths)
		os.Exit(1)
	}
	paths.Lemin(allPaths, graph.Nants)
}
