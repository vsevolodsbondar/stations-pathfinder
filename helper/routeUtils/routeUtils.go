package routeutils

import (
	"slices"
	m "trains/models"
)

func SortRoutesByDistance(routes []m.Route) []m.Route {
	slices.SortFunc(routes, func(a, b m.Route) int {
		if a.Distance < b.Distance {
			return -1
		}

		if a.Distance > b.Distance {
			return 1
		}

		return 0
	})

	return routes
}

func findCrossingRoutes(routes []m.Route) {
	for i := 0; i < len(routes); i++ {
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
}
