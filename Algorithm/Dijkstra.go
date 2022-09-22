package Algorithm

import (
	"math"
)

/*Parameters
GRAPH = The adjacency matrix that represents the graph.
STARTNODE = The node to start travelling from
ENDNODE = The node to travel to
*/
//Returns an array that displays the shortest path taken.
func FindShortestPath(graph [][]int, startNode int, endNode int) []int {
	var distancesToNodes = map[int]int{}
	var pathTaken = make([]int, 0)
	currentNodeConnections := graph[startNode]
	for i := range graph {
		if currentNodeConnections[i] == 0 {
			distancesToNodes[i] = math.MaxInt
		} else {
			distancesToNodes[i] = currentNodeConnections[i]
		}
	}
	distancesToNodes[startNode] = 0
	currentNode := startNode
	//loop through every node in the graph
	ctr := 0
	for ctr < len(graph) {
		currentNodeConnections = graph[currentNode]
		//loop through the current nodes connections
		for k := range currentNodeConnections {
			//if this nodes connection is shorter than the current map of distances, update
			if currentNodeConnections[k] > distancesToNodes[k] { //CHECK THIS
				distancesToNodes[k] = currentNodeConnections[k]
			}
		}
		pathTaken = append(pathTaken, currentNode)
		currentNode = GetShortestPathFromNode(currentNodeConnections)
	}
	return pathTaken
}

// returns index of shortest path from a given node
func GetShortestPathFromNode(slice []int) int {
	shortest := math.MaxInt
	for i := range slice {
		if (slice[i] < shortest) && slice[i] != 0 {
			shortest = i
		}
	}
	return shortest
}
