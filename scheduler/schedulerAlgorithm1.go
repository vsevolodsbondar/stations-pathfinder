package scheduler

import (
	"fmt"
	"math"
	h "trains/helper/schedulingUtils"
	m "trains/models"
)

func MoveTrains(routes []m.Route, trainsNumb int) {
	trains := h.TrainMaker(routes[0].Route[0], trainsNumb)

	finished := AssignRouteToTrain(trains, routes)
	freeStations := map[string]bool{}

	//filling map with stations
	for _, v := range routes {
		for _, r := range v.Route[1 : len(v.Route)-1] {
			freeStations[r] = true
		}
	}

	turn := 1
	for !allFinished(finished) {
		fmt.Println("Turn:", turn)
		plannedMoves := map[string]int{}
		planTurn(freeStations, trains, plannedMoves)
		executeTurn(freeStations, trains, finished)
		turn++
	}
}

func planTurn(freeStations map[string]bool, trains []m.Train, plannedMoves map[string]int) {
	for i := range trains {
		if trains[i].Finished {
			continue
		}

		if trains[i].Turn == len(trains[i].Route.Route)-1 {
			continue
		}

		nextStation := trains[i].Route.Route[trains[i].Turn+1] //sorry
		endStation := trains[i].Route.Route[len(trains[i].Route.Route)-1]

		if nextStation == endStation {
			trains[i].CanMove = true
			continue
		}

		if !freeStations[nextStation] {
			continue
		}

		_, v := plannedMoves[nextStation]
		if v {
			continue
		}

		trains[i].CanMove = true
		plannedMoves[nextStation] = trains[i].ID
	}
}

func executeTurn(freeStations map[string]bool, trains []m.Train, finished map[int]bool) {
	turnString := ""
	for i := range trains {
		if !trains[i].CanMove {
			continue
		}

		oldStation := trains[i].CurrStation

		moved, info := trains[i].Move()
		if !moved {
			finished[trains[i].ID] = true
		}

		newStation := trains[i].CurrStation
		endStation := trains[i].Route.Route[len(trains[i].Route.Route)-1]

		_, ok := freeStations[oldStation]
		if ok {
			freeStations[oldStation] = true
		}

		if newStation != endStation {
			freeStations[newStation] = false
		}

		turnString += info
	}

	if turnString != " " {
		fmt.Println(turnString)
	}

	turnString = ""
}

func allFinished(finished map[int]bool) bool {
	for _, done := range finished {
		if !done {
			return false
		}
	}

	return true
}

func AssignRouteToTrain(trains []m.Train, routes []m.Route) map[int]bool {
	finished := map[int]bool{}

	for i := range trains {
		route := bestCurrentRoute(routes)
		trains[i].Route = route

		finished[trains[i].ID] = false
	}

	return finished
}

func bestCurrentRoute(routes []m.Route) m.Route {
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

func BestSingleRoute(routeSets [][]m.Route) []m.Route {
	shortestRoute := []m.Route{}
	distance := math.MaxFloat64
	edges := math.MaxInt
	for i := 0; i < len(routeSets); i++ {
		for j := 0; j < len(routeSets[i]); j++ {
			if len(routeSets[i][j].Route) < edges {
				shortestRoute = append(shortestRoute, routeSets[i][j])
				distance = routeSets[i][j].Distance
				edges = len(routeSets[i][j].Route)
			} else if len(routeSets[i][j].Route) == edges {
				if routeSets[i][j].Distance < distance {
					shortestRoute = append(shortestRoute, routeSets[i][j])
					distance = routeSets[i][j].Distance
				}
			}
		}
	}

	return shortestRoute
}

func BestMultipleRoutes(routeSets [][]m.Route, trains int) []m.Route {
	mostEffectiveIndependentRoutes := []m.Route{}

	maxIndependentRoutes := 0
	for _, v := range routeSets {
		if len(v) > maxIndependentRoutes {
			maxIndependentRoutes = len(v)
		}
	}

	if trains < maxIndependentRoutes {
		maxIndependentRoutes = trains
	}

	sets := [][]m.Route{}

	for _, v := range routeSets {
		if len(v) == maxIndependentRoutes {
			sets = append(sets, v)
		}
	}

	if len(sets) == 1 {
		return sets[0]
	}

	currEdgesAvg := math.MaxFloat64
	for _, v := range sets {
		edges := 0
		counter := 0
		for _, r := range v {
			edges += len(r.Route)
			counter++
		}

		avg := float64(edges / counter)

		if avg < currEdgesAvg {
			mostEffectiveIndependentRoutes = append([]m.Route{}, v...)
			currEdgesAvg = float64(avg)
		}
	}

	return mostEffectiveIndependentRoutes
}
