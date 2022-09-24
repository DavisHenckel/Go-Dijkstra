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
func FindShortestPath(graph [][]int, startNode int, endNode int) *GraphNode {
	//initialize the priority queue using a min heap.
	prioQueue := make(PriorityQueue, 0)
	heap.Init(&prioQueue)
	var finalNode *GraphNode
	//hashmap to keep track of overall distances.
	var distancesToNodes = map[int]int{}
	currentNodeConnections := graph[startNode]
	//add all nodes to priority queue with int max as the Distance
	for i := range graph {
		if i == 0 {
			ithNode := GraphNode{Name: startNode, Previous: nil, Distance: 0, Index: i}
			heap.Push(&prioQueue, &ithNode)
			distancesToNodes[startNode] = 0
		} else {
			distancesToNodes[i] = math.MaxInt
			ithNode := GraphNode{Name: i, Previous: nil, Distance: math.MaxInt, Index: i}
			heap.Push(&prioQueue, &ithNode)
			distancesToNodes[i] = math.MaxInt
		}
	}
	//loop through every node in the graph
	for len(prioQueue) > 0 {
		currentNode := heap.Pop(&prioQueue).(*GraphNode)
		currentNodeConnections = graph[currentNode.Name]
		//loop through the current nodes connections
		for i := range currentNodeConnections {
			if currentNodeConnections[i] != 0 {
				// get the node in question
				var nodeToUpdate *GraphNode
				j := 0
				for j < len(prioQueue) {
					tmp := prioQueue[j]
					if tmp.Name == i {
						nodeToUpdate = tmp
						break
					}
					j++
				}
				//update the node only if the Distance has changed
				//if this nodes connection is shorter than the current map of distances, update the hashmap
				if currentNodeConnections[i]+currentNode.Distance < distancesToNodes[i] {
					distancesToNodes[i] = currentNodeConnections[i] + currentNode.Distance
					if nodeToUpdate != nil {
						prioQueue.update(nodeToUpdate, currentNodeConnections[i]+currentNode.Distance, currentNode)
						if endNode == nodeToUpdate.Name {
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

// returns Index of shortest path from a given node
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
