package AdjacencyMatrix

import "testing"

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
