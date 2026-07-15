package pathfinder

import (
	"math"
	m "trains/models"
)

// dfs + backtracking
// [][]string to just show Station.Name
func DFS(appData m.AppData) [][]string {
	res := [][]string{}
	visited := make(map[string]bool)

	// give to recursion full path that was already made
	// base case to check if it's end point
	// if not - for loop on all conections
	// in loop call of the recursion

	route := []string{appData.StartingStation.Name}
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

	recursion(route, appData.StartingStation)

	return res
}

func DFSRangeRoutes(appData m.AppData) map[int]m.Route {
	paths := DFS(appData)

	res := make(map[int]m.Route)

	for i, p := range paths {
		route := pathDistance(appData, p)

		res[i] = route
	}

	return res
}

func pathDistance(appData m.AppData, path []string) m.Route {
	var distance float64

	for i := 0; i < len(path)-1; i++ {
		from := appData.FindStationByName(path[i])
		to := appData.FindStationByName(path[i+1])

		distance += findDistance(from, to)
	}

	return m.Route{
		Route:    path,
		Distance: distance,
	}
}

func findDistance(station1, station2 *m.Station) float64 {
	diffX := float64(station1.X_axis - station2.X_axis)
	diffY := float64(station1.Y_axis - station2.Y_axis)

	distance := math.Sqrt(diffX*diffX + diffY*diffY)

	return distance
}
