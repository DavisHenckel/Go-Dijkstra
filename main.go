package main

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"Coding/Go/Go-Dijkstra/Algorithm"
	"fmt"
	"math"
)

func printNodeDistances(allNodes map[int]*Algorithm.GraphNode, startNode int) {
	fmt.Println("Total distances from", startNode, "to each node")
	i := 0
	for i < len(allNodes) {
		if allNodes[i].Distance == math.MaxInt {
			fmt.Println("Node", allNodes[i].Name, "-- Path does not exist")
			i++
		}
		fmt.Println("Node:", allNodes[i].Name, "-- Distance:", allNodes[i].Distance)
		i++
	}
}

func printResult(finalNode Algorithm.GraphNode, startNode int, endNode int) {
	if finalNode.Distance == math.MaxInt {
		fmt.Println("There is no path from node", startNode, "to node", endNode)
		return
	}
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

func getNumNodes() int {
	var numNodes int
	fmt.Print("Enter the number of nodes for this graph: ")
	_, err := fmt.Scan(&numNodes)
	for err != nil {
		fmt.Println("You did not enter a number. Picking a number at random...")
		numNodes = AdjacencyMatrix.PickNumNodes()
		break
	}
	if numNodes < 9 || numNodes > 1000 {
		fmt.Println("Nodes must be at least 10 and no more than 1000. Selecting 10 for now...")
		numNodes = 10
	}
	return numNodes
}

func printMenu() int {
	var choice int
	for choice < 1 || choice > 5 {
		fmt.Println("\nPlease choose from the following options")
		fmt.Println("1: Print Adjacency Matrix")
		fmt.Println("2: Print Raw Matrix (easy copy/paste)")
		fmt.Println("3: Find Shortest Path from Node A to B")
		fmt.Println("4: Find Shortest Path to all Nodes")
		fmt.Println("5: Create New Adjacency Matrix")
		fmt.Scan(&choice)
	}
	return choice
}

func pickEndNode(graphSize int) int {
	var nodeB int
	fmt.Print("Pick destination node: ")
	fmt.Scan(&nodeB)
	for nodeB > graphSize-1 || nodeB < 0 {
		fmt.Print("Node: ", nodeB, " doesn't exist in this graph\nPick desintation node: ")
		fmt.Scan(&nodeB)
	}
	return nodeB
}

func pickStartNode(graphSize int) int {
	var nodeA int
	fmt.Print("Pick starting node: ")
	fmt.Scan(&nodeA)
	for nodeA > graphSize-1 || nodeA < 0 {
		fmt.Print("Node ", nodeA, " doesn't exist in this graph\nPick starting node: ")
		fmt.Scan(&nodeA)
	}
	return nodeA
}

func printSpecificNodePath(nodes map[int]*Algorithm.GraphNode, startNode int) {
	var nodeChosen int
	i := 0
	for i < 2 {
		fmt.Print("Enter any node to see its path, enter any key to return to the menu: ")
		_, err := fmt.Scan(&nodeChosen)
		if err != nil {
			return
		}
		if nodeChosen > len(nodes)-1 || nodeChosen < 0 {
			continue
		} else {
			printResult(*nodes[nodeChosen], startNode, nodeChosen)
		}
	}

}

func main() {
	fmt.Println("Welcome to Davis Henckel's implementation of Dijkstra's Algorithm written in Go.")
	fmt.Println("================================================================================")
	numNodes := getNumNodes()
	Graph := AdjacencyMatrix.GenerateMatrix(numNodes)
	userInput := printMenu()
	i := 0
	for i < 2 {
		switch userInput {
		case 1:
			AdjacencyMatrix.PrintMatrix(Graph, len(Graph))
			userInput = printMenu()
		case 2:
			AdjacencyMatrix.PrintRawMatrix(Graph, ",")
			userInput = printMenu()
		case 3:
			startNode := pickStartNode(len(Graph))
			endNode := pickEndNode(len(Graph))
			finalNode, _ := Algorithm.FindShortestPath(Graph, startNode, endNode)
			if finalNode == nil {
				fmt.Println("There is no path from node", startNode, "to node", endNode)
			} else {
				printResult(*finalNode, startNode, endNode)
			}
			userInput = printMenu()
		case 4:
			startNode := pickStartNode(len(Graph))
			_, allNodes := Algorithm.FindShortestPath(Graph, startNode, 10)
			printNodeDistances(allNodes, startNode)
			printSpecificNodePath(allNodes, startNode)
			userInput = printMenu()
		case 5:
			numNodes := getNumNodes()
			fmt.Println("\nThe number of nodes in this graph is", numNodes)
			Graph = AdjacencyMatrix.GenerateMatrix(numNodes)
			userInput = printMenu()
		}
	}
}
