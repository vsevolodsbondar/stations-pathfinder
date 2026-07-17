package pathfinder

import (
	"fmt"
	d "trains/helper/distanceCalculations"
	h "trains/helper/routeUtils"
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

func DFSRangedRouteSets(appData m.AppData) ([][]m.Route, error) {
	paths, err := DFS(appData)
	if err != nil {
		return nil, err
	}

	res := []m.Route{}

	for i, p := range paths {
		route := d.PathsDistance(appData, p)
		route.ID = i + 1
		route.CrossingRoutes = map[int]struct{}{}

		res = append(res, route)
	}

	m := h.FindUniqueRouteSets(res)

	return m, nil
}
