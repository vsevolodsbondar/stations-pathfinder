package scheduler

import (
	"fmt"
	h "trains/helper/schedulingUtils"
	m "trains/models"
)

func MoveTrains(routes []m.Route, trainsNumb int) {
	trains := h.TrainMaker(routes[0].Route[0], trainsNumb)
	finished := AssignRouteToTrain(trains, routes)

	freeStations := map[string]bool{}

	// Only intermediate stations are tracked.
	for _, r := range routes {
		for _, s := range r.Route[1 : len(r.Route)-1] {
			freeStations[s] = true
		}
	}

	for !allFinished(finished) {
		turnString := ""

		for i := range trains {
			if trains[i].Finished {
				continue
			}

			nextStation := trains[i].Route.Route[trains[i].Turn+1]
			endStation := trains[i].Route.Route[len(trains[i].Route.Route)-1]

			if nextStation != endStation && !freeStations[nextStation] {
				continue
			}

			prevStation := trains[i].CurrStation

			// freeing prev station
			_, exists := freeStations[prevStation]
			if exists {
				freeStations[prevStation] = true
			}

			// occupying next station if it's not the last one
			if nextStation != endStation {
				freeStations[nextStation] = false
			}

			moved, info := trains[i].Move()
			turnString += info

			if !moved {
				finished[trains[i].ID] = true
			}
		}

		if turnString != "" {
			fmt.Println(turnString)
		}
	}
}

func AssignRouteToTrain(trains []m.Train, routes []m.Route) map[int]bool {
	finished := map[int]bool{}

	for i := range trains {
		route := bestTrainRoute(routes)
		trains[i].Route = route

		finished[trains[i].ID] = false
	}

	return finished
}

func bestTrainRoute(routes []m.Route) m.Route {
	best := 0

	for i := 1; i < len(routes); i++ {
		if routes[i].RouteScore < routes[best].RouteScore {
			best = i
		} else if routes[i].RouteScore == routes[best].RouteScore &&
			routes[i].Distance < routes[best].Distance {
			best = i
		}
	}

	routes[best].RouteScore++

	return routes[best]
}

func allFinished(finished map[int]bool) bool {
	for _, done := range finished {
		if !done {
			return false
		}
	}

	return true
}
