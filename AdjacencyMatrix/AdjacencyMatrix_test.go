package AdjacencyMatrix

import "testing"

func TestMatrixSize(t *testing.T) {
	i := 0
	for i < 10000 {
		num := PickMatrixSize()
		if num%2 != 0 || num > 50 || num < 16 {
			t.Fatalf(`PickMatrixSize returned %v`, num)
		}
		i++
	}
}
