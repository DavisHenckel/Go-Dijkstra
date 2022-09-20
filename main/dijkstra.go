package main

import (
	"Coding/Go/Go-Dijkstra/AdjacencyMatrix"
	"fmt"
)

func main() {
	fmt.Println(AdjacencyMatrix.PickNumNodes())
	AdjacencyMatrix.GenerateMatrix(AdjacencyMatrix.PickNumNodes())
}
