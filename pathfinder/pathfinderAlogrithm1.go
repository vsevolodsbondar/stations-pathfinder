package pathfinder

import (
	"fmt"
	"math"
	"slices"
	m "trains/models"
)

// dfs + backtracking
func DFS(appData m.AppData) ([][]string, error) {
	res := [][]string{}
	visited := make(map[string]bool)

	// give to recursion full path that was already made
	// base case to check if it's end point
	// if not - for loop on all conections
	// in loop call of the recursion

	startingRoute := []string{appData.StartingStation.Name}
	var recursion func([]string, *m.Station)

	recursion = func(route []string, current *m.Station) {
		//base case
		if current.Name == appData.EndingStation.Name {
			res = append(res, route)
			return
		}

		//will be checking if I already checked all connections of particular station
		//helps to avoid loops (already had to reboot PC once)
		visited[current.Name] = true
		for _, v := range current.Connections {
			copyOfPrevRoute := append([]string{}, route...)
			newRoute := append(copyOfPrevRoute, v.Name)

			//stoping here if visited
			if visited[v.Name] {
				continue
			}

			recursion(newRoute, v)
		}

		//checked all connections for the station, unblocking
		visited[current.Name] = false
	}

	recursion(startingRoute, appData.StartingStation)

	if len(res) == 0 {
		return nil, fmt.Errorf("There is no route that reaches ending station.")
	}
	return res, nil
}

func DFSRangeRoutes(appData m.AppData) ([]m.Route, error) {
	paths, err := DFS(appData)
	if err != nil {
		return nil, err
	}

	res := []m.Route{}

	for _, p := range paths {
		route := pathsDistance(appData, p)

		res = append(res, route)
	}

	res = sortRoutesByDistance(res)

	return res, nil
}

func pathsDistance(appData m.AppData, path []string) m.Route {
	var distance float64

	for i := 0; i < len(path)-1; i++ {
		from := appData.FindStationByName(path[i])
		to := appData.FindStationByName(path[i+1])

		distance += findDistanceBetweenPoints(from, to)
	}

	return m.Route{
		Route:    path,
		Distance: distance,
	}
}

func findDistanceBetweenPoints(station1, station2 *m.Station) float64 {
	diffX := float64(station1.X_axis - station2.X_axis)
	diffY := float64(station1.Y_axis - station2.Y_axis)

	distance := math.Sqrt(diffX*diffX + diffY*diffY)

	return distance
}

func sortRoutesByDistance(routes []m.Route) []m.Route {
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

// func findCrossingRoutes(routes []m.Route) []m.Route {

// }
