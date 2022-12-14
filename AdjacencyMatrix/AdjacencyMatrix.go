package AdjacencyMatrix

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns an even number, with a minimum of 15 and a maximum of 75
func PickNumNodes() int {
	returnValue := rand.Intn(60) + 15
	return returnValue
}

// Return a 2d slice that represents the adjacency matrix
func GenerateMatrix(size int) [][]int {
	adjacencyMatrix := make([][]int, size)
	for i := range adjacencyMatrix {
		if adjacencyMatrix[i] == nil {
			adjacencyMatrix[i] = make([]int, size)
		}
		vertexDegree := rand.Intn(3) + 1 //pick how many connections this node will have
		for j := countVertexDegree(adjacencyMatrix[i]); j < vertexDegree; j++ {
			connection := rand.Intn(size)
			//ensure connecteed node is not exceeding the possible 4 degree
			if countVertexDegree(adjacencyMatrix[connection]) >= 4 {
				j--
				continue
			}
			//node cannot be connected to itself
			if connection == i {
				j--
				continue
			}
			//if there is already a connection
			if adjacencyMatrix[i][connection] != 0 {
				continue
			}
			distance := rand.Intn(8) + 1
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

func countVertexDegree(vertexSlice []int) int {
	var numConnections int
	for i := range vertexSlice {
		if vertexSlice[i] != 0 {
			numConnections++
		}
	}
	return numConnections
}

func PrintRawMatrix(matrix [][]int, delimiter string) {
	if len(delimiter) != 1 {
		fmt.Println("delimiter must be of length 1. Setting to ',' for now")
		delimiter = ","
	}
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix); y++ {
			if y == len(matrix)-1 {
				fmt.Println(matrix[x][y])
			} else {
				fmt.Print(matrix[x][y], delimiter)
			}
		}
	}
}

func PrintMatrix(matrix [][]int, size int) {
	PrintMatrixBorder(size)
	fmt.Print("| ")
	for x := 0; x < size; x++ {
		if x < 9 {
			fmt.Print(x, "  ")
		} else {
			if x == size-1 {
				fmt.Println(x, "|#####|")
			} else {
				fmt.Print(x, " ")
			}
		}
	}
	PrintMatrixBorder(size)
	counter := 0
	for i := range matrix {
		fmt.Print("| ")
		for j := range matrix[i] {
			if j == size-1 {
				fmt.Print(matrix[i][j])
			} else {
				fmt.Print(matrix[i][j], "  ")
			}
		}
		if counter < 10 {
			fmt.Println(" | ", counter, " |")
		} else {
			fmt.Println(" | ", counter, "|")
		}

		counter++
	}
	PrintMatrixBorder(size)
}

func PrintMatrixBorder(size int) {
	for x := 0; x < size*3+8; x++ {
		if x == size*3+7 {
			fmt.Println("-")
		} else {
			fmt.Print("-")
		}
	}
}
