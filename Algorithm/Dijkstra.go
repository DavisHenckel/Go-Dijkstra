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
	var pathTaken = make([]int, 1)
	startingSlice := graph[startNode]
	for i := range graph {
		if startingSlice[i] == 0 {
			distancesToNodes[i] = math.MaxInt
		} else {
			distancesToNodes[i] = startingSlice[i]
		}
	}
	distancesToNodes[startNode] = 0

	// loop through every node in the graph
	for ctr := range graph {

	}

	return pathTaken
}
