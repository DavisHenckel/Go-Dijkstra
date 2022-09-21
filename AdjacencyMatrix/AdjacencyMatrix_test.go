package AdjacencyMatrix

import (
	"math/rand"
	"testing"
)

func TestPickNumNodes(t *testing.T) {
	i := 0
	for i < 10000 {
		num := PickNumNodes()
		if num%2 != 0 || num > 50 || num < 16 {
			t.Fatalf(`PickNumNodes returned %v`, num)
		}
		i++
	}
}

func TestGenerateMatrixDegree(t *testing.T) {
	k := 0
	for k < 1000 {
		numNodes := PickNumNodes()
		matrix := GenerateMatrix(numNodes)
		i := 0
		for i < numNodes {
			if countVertexDegree(matrix[i]) > 4 {
				PrintMatrix(matrix, numNodes)
				vertDeg := countVertexDegree(matrix[i])
				t.Fatalf(`vertex %v has %v which exceeds the possible 4!`, i, vertDeg)
			}
			i++
		}
		k++
	}
}

func TestGenerateMatrix(t *testing.T) {
	numNodes := PickNumNodes()
	matrix := GenerateMatrix(numNodes)
	i := 0
	for i < numNodes {
		if matrix[i][i] != 0 {
			t.Fatalf(`Vertex %v does not contain 0 as it should. This matrix is not symmetric!`, i)
		}
		i++
	}
}

func TestGenerateMatrixSymmetry(t *testing.T) {
	numNodes := PickNumNodes()
	matrix := GenerateMatrix(numNodes)
	i := 0
	for i < numNodes {
		j := 0
		for j < numNodes {
			if matrix[i][j] != matrix[j][i] {
				PrintMatrix(matrix, numNodes)
				t.Fatalf(`matrix is not symmetrix at %v,%v`, i, j)
			}
			j++
		}
		i++
	}
}

func TestCountVertexDegree(t *testing.T) {
	length := 10
	numTrials := 0
	for numTrials < 100 {
		degree := rand.Intn(9)
		slice := make([]int, length)
		i := 0
		for i < degree {
			slice[i] = 1
			i++
		}
		result := countVertexDegree(slice)
		if result != degree {
			t.Fatalf(`degree does not match. Function returned %v while it should have been %v`, result, degree)
		}
		numTrials++
	}
}
