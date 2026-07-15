package pathfinder

import (
	m "trains/models"
)

// dfs + backtracking
// [][]string to just show Station.Name
func BigFuckingSearch(appData m.AppDataPointer) [][]string {
	res := [][]string{}
	visited := make(map[string]bool)

	// give to recursion full path that was already made
	// basic case to check if it's end point
	// if not - for loop on all conections
	// in loop call of the recursion

	path := []string{appData.StartingStation.Name}
	var recursion func([]string, m.PointerStation)

	recursion = func(path []string, current m.PointerStation) {
		if current.Name == appData.EndingStation.Name {
			res = append(res, path)
			return
		}

		visited[current.Name] = true
		for _, v := range current.Connections {
			newPath := append(append([]string{}, path...), v.Name)

			//stoping here if visited
			if visited[v.Name] {
				continue
			}

			recursion(newPath, *v)
		}

		visited[current.Name] = false
	}

	recursion(path, appData.StartingStation)

	return res
}
