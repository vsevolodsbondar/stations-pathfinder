package routeutils

import (
	"math"
	"slices"
	m "trains/models"
)

func sortRoutesByDistance(routes []m.Route) []m.Route {
	slices.SortFunc(routes, func(a, b m.Route) int {
		if a.Distance < b.Distance {
			return 1
		}

		if a.Distance > b.Distance {
			return -1
		}

		return 0
	})

	return routes
}

// for determining if there some overlaping of the routes
func findCrossingRoutes(routes []m.Route) []m.Route {
	for i, _ := range routes {
		for j := i + 1; j < len(routes); j++ {
			routeStations1 := routes[i].Route[1 : len(routes[i].Route)-1]
			routeStations2 := routes[j].Route[1 : len(routes[j].Route)-1]

			for _, station := range routeStations1 {
				if slices.Contains(routeStations2, station) {
					routes[i].CrossingRoutes[routes[j].ID] = struct{}{}
					routes[j].CrossingRoutes[routes[i].ID] = struct{}{}
					break
				}
			}
		}
	}
	return routes
}

func FindUniqueRouteSets(routes []m.Route) [][]m.Route {
	routes = findCrossingRoutes(sortRoutesByDistance(routes))
	var result [][]m.Route
	var current []m.Route

	var recursion func(routeIndx int)

	recursion = func(routeIndx int) {
		if routeIndx == len(routes) { // means I've checked all routes
			maximal := true

			//checking every route that is not in current []m.Route already
			for _, candidate := range routes {
				inCurrent := false
				for _, r := range current {
					if r.ID == candidate.ID {
						inCurrent = true
					}
				}

				if inCurrent == true {
					continue
				}

				canAdd := true
				for _, chosen := range current {
					_, crosses := chosen.CrossingRoutes[candidate.ID]
					if crosses {
						canAdd = false
						break
					}
				}

				if canAdd {
					maximal = false
					break
				}
			}

			if maximal {
				set := append([]m.Route{}, current...)
				result = append(result, set)
			}
			return
		}

		//skip this route
		recursion(routeIndx + 1)

		canAdd := true
		//checking if already added routes in current []m.Route conflict with current route
		for _, r := range current {
			_, crosses := routes[routeIndx].CrossingRoutes[r.ID]
			if crosses {
				canAdd = false
				break
			}
		}

		//take this route if possible
		if canAdd {
			current = append(current, routes[routeIndx])
			recursion(routeIndx + 1)
			current = current[:len(current)-1]
		}
	}

	recursion(0)

	return result
}

func BestRoutes(routeSets [][]m.Route, trains int) []m.Route {
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

// not used currently
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
