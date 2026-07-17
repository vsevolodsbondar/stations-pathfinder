package routeutils

import (
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
