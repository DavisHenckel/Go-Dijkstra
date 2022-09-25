package Algorithm

import (
	"container/heap"
)

//example code taken and modified from https://pkg.go.dev/container/heap@go1.19.1

type GraphNode struct {
	Name     int
	Previous *GraphNode
	Distance int
	Index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*GraphNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Want Pop to give us the lowest, not highest, Distance so use greater than here.
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*GraphNode)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the Distance and Previous of an Item in the queue.
func (pq *PriorityQueue) update(item *GraphNode, Distance int, Previous *GraphNode) {
	item.Distance = Distance
	item.Previous = Previous
	heap.Fix(pq, item.Index)
}
