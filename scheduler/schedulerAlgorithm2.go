package scheduler

import (
	"fmt"
	m "trains/models"
)

type Train struct {
	ID       int
	PathID   int
	Position int
}

// DistributeTrains assigns trains to the available paths so that the estimated
// arrival time of the last train is minimized. Shorter paths receive more trains
// because they become available again sooner.
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

// LaunchTrains starts new trains on their assigned paths whenever the first
// station after the start is free. Newly launched trains are added to the list
// of active trains and receive consecutive IDs.
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

// MoveTrains advances all active trains by one station if their next station
// is available. It updates occupied stations, records train movements for the
// current turn, and removes trains that have reached the destination.
func MoveActiveTrains(
	active *[]Train,
	paths [][]*m.Station,
	occupied map[*m.Station]bool,
) []string {

	for k := range occupied {
		delete(occupied, k)
	}

	for _, train := range *active {
		path := paths[train.PathID]
		if train.Position < len(path)-1 {
			cur := path[train.Position]
			if cur != path[len(path)-1] {
				occupied[cur] = true
			}
		}
	}

	moves := []string{}

	for i := range *active {

		train := &(*active)[i]
		path := paths[train.PathID]

		if train.Position >= len(path)-1 {
			continue
		}

		current := path[train.Position]
		next := path[train.Position+1]

		if current != path[len(path)-1] {
			delete(occupied, current)
		}

		if next != path[len(path)-1] {
			if occupied[next] {
				if current != path[len(path)-1] {
					occupied[current] = true
				}
				continue
			}
			occupied[next] = true
		}

		train.Position++
		moves = append(moves, fmt.Sprintf("T%d-%s", train.ID, next.Name))
	}

	newActive := (*active)[:0]

	for _, train := range *active {
		if train.Position < len(paths[train.PathID])-1 {
			newActive = append(newActive, train)
		}
	}

	*active = newActive

	return moves
}
