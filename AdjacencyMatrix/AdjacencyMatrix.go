package AdjacencyMatrix

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns an even number, with a minimum of 16 and a maximum of 50
func PickNumNodes() int {
	returnValue := rand.Intn(34) + 16
	for returnValue%2 != 0 {
		returnValue = rand.Intn(34) + 16
	}
	return returnValue
}

// Return a 2d slice that represents the adjacency matrix
func GenerateMatrix(size int) [][]int {
	adjacencyMatrix := make([][]int, size)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, size)
		numConnections := rand.Intn(5) + 1 //pick how many connections this node will have
		for j := 0; j < numConnections; j++ {
			connection := rand.Intn(size)
			//node cannot be connected to itself
			if connection == i {
				j--
				continue
			}
			//if there is already a connection
			if adjacencyMatrix[i][connection] != 0 {
				continue
			}
			distance := rand.Intn(9) + 1
			adjacencyMatrix[i][connection] = distance
			//assign the other node to also be connected to i node
			if adjacencyMatrix[connection] == nil {
				adjacencyMatrix[connection] = make([]int, size)
				adjacencyMatrix[connection][i] = distance
			} else {
				adjacencyMatrix[connection][i] = distance
			}

		}
	}
	return adjacencyMatrix
}
