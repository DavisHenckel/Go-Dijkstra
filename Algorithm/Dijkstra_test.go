package Algorithm

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	FindShortestPath(AdjacencyMatrix.GenerateMatrix(AdjacencyMatrix.PickNumNodes()), 1, 8)
}
