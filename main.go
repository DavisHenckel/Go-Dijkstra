package main

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"Coding/Go/Go-Dijkstra/Algorithm"
	"fmt"
)

// graph for testing. There are 3 separate graphs
var disconnectedGraph = [][]int{
	{0, 0, 0, 5, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0},
	{0, 0, 8, 0, 0, 0, 3, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 5, 0, 0},
	{0, 0, 0, 0, 0, 0, 5, 4, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 3, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 6, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 5, 0, 6, 0, 0, 0},
	{0, 4, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
	{6, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0},
}

func printNodeDistances(allNodes map[int]*Algorithm.GraphNode, startNode int) {
	fmt.Println("Total distances from", startNode, "to each node")
	i := 0
	for i < len(allNodes) {
		fmt.Println("Node:", allNodes[i].Name, "-- Distance:", allNodes[i].Distance)
		i++
	}
}

func printResult(finalNode Algorithm.GraphNode, startNode int, endNode int) {
	fmt.Println("Shortest distance from node", startNode, "to node", endNode, "is", finalNode.Distance)
	fmt.Print("The path is: ")
	pathTaken, error := Algorithm.GetNodesPathTraversed(&finalNode)
	if error == nil {
		for i := range pathTaken {
			if i == len(pathTaken)-1 {
				fmt.Println(pathTaken[i])
			} else {
				fmt.Print(pathTaken[i], "->")
			}
		}
	}
}

func main() {
	startNode := 1
	endNode := 10
	fmt.Println("Welcome to Davis Henckel's implementation of Dijkstra's Algorithm written in Go.\n================================================================================")
	numNodes := AdjacencyMatrix.PickNumNodes()
	fmt.Println("\nThe number of nodes in this graph is", numNodes)
	Graph := AdjacencyMatrix.GenerateMatrix(numNodes)
	AdjacencyMatrix.PrintMatrix(Graph, numNodes)
	finalNode, allNodes := Algorithm.FindShortestPath(Graph, startNode, endNode)
	printResult(*finalNode, startNode, endNode)
	printNodeDistances(allNodes, startNode)
}
