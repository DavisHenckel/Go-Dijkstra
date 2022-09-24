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
func FindShortestPath(graph [][]int, startNode int, endNode int) *graphNode {
	//initialize the priority queue using a min heap.
	prioQueue := make(PriorityQueue, 0)
	heap.Init(&prioQueue)
	var finalNode *graphNode
	//hashmap to keep track of overall distances.
	var distancesToNodes = map[int]int{}
	currentNodeConnections := graph[startNode]
	//add all nodes to priority queue with int max as the distance
	for i := range graph {
		if i == 0 {
			ithNode := graphNode{name: startNode, previous: nil, distance: 0, index: i}
			heap.Push(&prioQueue, &ithNode)
			distancesToNodes[startNode] = 0
		} else {
			distancesToNodes[i] = math.MaxInt
			ithNode := graphNode{name: i, previous: nil, distance: math.MaxInt, index: i}
			heap.Push(&prioQueue, &ithNode)
			distancesToNodes[i] = math.MaxInt
		}
	}
	//loop through every node in the graph
	for len(prioQueue) > 0 {
		currentNode := heap.Pop(&prioQueue).(*graphNode)
		currentNodeConnections = graph[currentNode.name]
		//loop through the current nodes connections
		for i := range currentNodeConnections {
			if currentNodeConnections[i] != 0 {
				// get the node in question
				var nodeToUpdate *graphNode
				j := 0
				for j < len(prioQueue) {
					tmp := prioQueue[j]
					if tmp.name == i {
						nodeToUpdate = tmp
						break
					}
					j++
				}
				//update the node only if the distance has changed
				//if this nodes connection is shorter than the current map of distances, update the hashmap
				if currentNodeConnections[i]+currentNode.distance < distancesToNodes[i] {
					distancesToNodes[i] = currentNodeConnections[i] + currentNode.distance
					if nodeToUpdate != nil {
						prioQueue.update(nodeToUpdate, currentNodeConnections[i]+currentNode.distance, currentNode)
						if endNode == nodeToUpdate.name {
							finalNode = nodeToUpdate
						}
					}
				}
			}
		}
	}
	return finalNode
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
