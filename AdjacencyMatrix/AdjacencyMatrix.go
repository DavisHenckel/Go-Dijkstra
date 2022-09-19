package AdjacencyMatrix

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns an even number, with a minimum of 16 and a maximum of 50
func PickMatrixSize() int {
	returnValue := rand.Intn(34) + 16
	for returnValue%2 != 0 {
		returnValue = rand.Intn(34) + 16
	}
	return returnValue
}
