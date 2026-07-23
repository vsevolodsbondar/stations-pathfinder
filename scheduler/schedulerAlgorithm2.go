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
	cost := make([]int, len(paths))

	for i, path := range paths {
		cost[i] = len(path)
	}

	for i := 0; i < trainCount; i++ {
		best := 0

		for j := 1; j < len(cost); j++ {
			if cost[j] < cost[best] {
				best = j
			}
		}

		assigned[best]++
		cost[best]++
	}
	return assigned
}
