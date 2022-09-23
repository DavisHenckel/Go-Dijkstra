package Algorithm

import (
	"container/heap"
)

//example code taken and modified from https://pkg.go.dev/container/heap@go1.19.1

type graphNode struct {
	name     int
	previous *graphNode
	distance int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*graphNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Want Pop to give us the lowest, not highest, distance so use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*graphNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the distance and name of an Item in the queue.
func (pq *PriorityQueue) update(item *graphNode, name int, distance int) {
	item.name = name
	item.distance = distance
	heap.Fix(pq, item.index)
}
