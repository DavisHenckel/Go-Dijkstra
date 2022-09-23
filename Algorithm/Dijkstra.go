package Algorithm

import (
	"container/heap"
	"math"
)

/*Parameters
GRAPH = The adjacency matrix that represents the graph.
STARTNODE = The node to start travelling from
ENDNODE = The node to travel to
*/
//Returns an array that displays the shortest path taken.
func FindShortestPath(graph [][]int, startNode int, endNode int) []int {
	//initialize the priority queue
	prioQueue := make(PriorityQueue, 0)
	// currentNode := graphNode{name: startNode, previous: nil, distance: 10, index: 0}
	// currentNode1 := graphNode{name: 10, previous: &currentNode, distance: 2, index: 1}
	heap.Init(&prioQueue)
	// heap.Push(&prioQueue, &currentNode)
	// heap.Push(&prioQueue, currentNode)
	//hashmap to keep track of overall distances.
	var distancesToNodes = map[int]int{}
	var nodesVisited = make([]int, 0)
	var nodesUnvisited = make([]int, 0)
	currentNodeConnections := graph[startNode]
	for i := range graph {
		//add all nodes to nodesUnvisited
		nodesUnvisited = append(nodesUnvisited, i)
		//setup for current node
		if currentNodeConnections[i] == 0 {
			distancesToNodes[i] = math.MaxInt
		} else {
			distancesToNodes[i] = currentNodeConnections[i]
		}
	}
	distancesToNodes[startNode] = 0
	currentNode := graphNode{name: startNode, previous: nil, distance: 0}
	nodesUnvisited = removeElement(nodesUnvisited, startNode)
	//loop through every node in the graph
	for len(nodesVisited) != len(graph) {
		currentNodeConnections = graph[currentNode.name]
		//loop through the current nodes connections
		for i := range currentNodeConnections {
			var nodeToAdd graphNode
			if currentNodeConnections[i] != 0 {
				nodeToAdd = graphNode{name: i, previous: &currentNode, distance: currentNodeConnections[i]}
				heap.Push(&prioQueue, &nodeToAdd)
			}
			//if this nodes connection is shorter than the current map of distances, update
			if currentNodeConnections[i] < distancesToNodes[i] && currentNodeConnections[i] != 0 {
				//do something with heap here
				distancesToNodes[i] = currentNodeConnections[i]
			}
		}
		nodesVisited = append(nodesVisited, currentNode.name)
		nodesUnvisited = removeElement(nodesUnvisited, currentNode.name)
		newNode := heap.Pop(&prioQueue).(*graphNode)
		currentNode = *newNode
		//currentNode = GetShortestPathFromNode(currentNodeConnections, nodesVisited)
	}
	return nodesVisited
}

func removeElement(slice []int, element int) []int {
	var returnSlice = make([]int, 0)

	for i := range slice {
		if slice[i] == element {
			continue
		}
		returnSlice = append(returnSlice, slice[i])
	}
	return append(returnSlice)
}

// returns index of shortest path from a given node
func GetShortestPathFromNode(slice []int, nodesVisited []int) int {
	shortest := math.MaxInt
	for i := range slice {
		if ((slice[i] < shortest) && slice[i] != 0) && containsElement(nodesVisited, i) == false {
			shortest = i
		}
	}
	return shortest
}

func containsElement(slice []int, element int) bool {
	for i := range slice {
		if slice[i] == element {
			return true
		}
	}
	return false
}
