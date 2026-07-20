package scheduler

import (
	"math"
	m "trains/models"
)

func BestSingleRoute(routeSets [][]m.Route) m.Route {
	shortestRoute := m.Route{}
	distance := math.MaxFloat64
	edges := math.MaxInt
	for i := 0; i < len(routeSets); i++ {
		for j := 0; j < len(routeSets[i]); j++ {
			if len(routeSets[i][j].Route) < edges {
				shortestRoute = routeSets[i][j]
				distance = routeSets[i][j].Distance
				edges = len(routeSets[i][j].Route)
			} else if len(routeSets[i][j].Route) == edges {
				if routeSets[i][j].Distance < distance {
					shortestRoute = routeSets[i][j]
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
