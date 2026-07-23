package scheduler

import (
	m "trains/models"
)

type Train struct {
	ID       int
	PathID   int
	Position int
}

func DistributeTrains(paths [][]*m.Station, trainCount int) []int {
	assigned := make([]int, len(paths))
	arrival := make([]int, len(paths))

	for i, path := range paths {
		arrival[i] = len(path)
	}

	for i := 0; i < trainCount; i++ {
		best := 0

		for j := 1; j < len(arrival); j++ {
			if arrival[j] < arrival[best] {
				best = j
			}
		}

		assigned[best]++
		arrival[best]++
	}
	return assigned
}
