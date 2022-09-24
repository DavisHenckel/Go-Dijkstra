package main

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"Coding/Go/Go-Dijkstra/Algorithm"
	"fmt"
)

func main() {
	startNode := 1
	endNode := 10
	fmt.Println("Welcome to Davis Henckel's implementation of Dijkstra's Algorithm written in Go.\n================================================================================")
	numNodes := AdjacencyMatrix.PickNumNodes()
	fmt.Println("\nThe number of nodes in this graph is", numNodes)
	Graph := AdjacencyMatrix.GenerateMatrix(numNodes)
	AdjacencyMatrix.PrintMatrix(Graph, numNodes)
	finalNode, allNodes := Algorithm.FindShortestPath(Graph, startNode, endNode)
	pathTaken, error := Algorithm.GetNodesPathTraversed(finalNode)
	if error == nil {
		for i := range pathTaken {
			if i == len(pathTaken)-1 {
				fmt.Println(pathTaken[i])
			} else {
				fmt.Print(pathTaken[i], "->")
			}
		}
	}
	i := 0
	for i < len(allNodes) {
		t := allNodes[i]
		fmt.Println(t.Name, " -- ", allNodes[i].Distance)
		i++
	}

}
