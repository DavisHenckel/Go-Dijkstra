package main

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"Coding/Go/Go-Dijkstra/Algorithm"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Davis Henckel's implementation of Dijkstra's Algorithm written in Go.\n================================================================================")
	numNodes := AdjacencyMatrix.PickNumNodes()
	fmt.Println("\nThe number of nodes in this graph is", numNodes)
	Graph := AdjacencyMatrix.GenerateMatrix(numNodes)
	AdjacencyMatrix.PrintMatrix(Graph, numNodes)
	finalNode := Algorithm.FindShortestPath(Graph, 1, 10)
	fmt.Println("Found the shortest path to node ", finalNode.Name, ", the distance was ", finalNode.Distance)
	pathTaken := make([]int, 0)
	pathTaken = append(pathTaken, finalNode.Name)
	tmp := finalNode.Previous
	for tmp != nil {
		pathTaken = append(pathTaken, tmp.Name)
		tmp = tmp.Previous
	}
	fmt.Print("The path was: ")
	i := len(pathTaken) - 1
	for i >= 0 {
		if i == 0 {
			fmt.Println(pathTaken[i])
		} else {
			fmt.Print(pathTaken[i], "->")
		}
		i--
	}
}
