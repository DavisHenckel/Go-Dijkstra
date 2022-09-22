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
	Algorithm.Hello()

}
