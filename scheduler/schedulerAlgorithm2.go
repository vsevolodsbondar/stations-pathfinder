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

func LaunchTrains(
	active *[]Train,
	paths [][]*m.Station,
	assigned []int,
	launched []int,
	occupied map[*m.Station]bool,
	nextID *int,
) {

	for pathID, path := range paths {

		if launched[pathID] >= assigned[pathID] {
			continue
		}

		if len(path) < 2 {
			continue
		}

		first := path[1]

		if first != path[len(path)-1] && occupied[first] {
			continue
		}

		*active = append(*active, Train{
			ID:       *nextID,
			PathID:   pathID,
			Position: 1,
		})

		if first != path[len(path)-1] {
			occupied[first] = true
		}

		launched[pathID]++
		(*nextID)++
	}
}
