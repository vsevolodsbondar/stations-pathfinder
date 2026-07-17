package pathfinder
<<<<<<< HEAD
=======

import (
	m "trains/models"
)

func BFS(appData m.AppData) []string {
	var visited map[string]struct{}
	res := []string{}
	start := appData.StartingStation.Name
	end := appData.EndingStation.Name
	posMax := 0
	if len(appData.StartingStation.Connections) > len(appData.EndingStation.Connections) {
		posMax = len(appData.EndingStation.Connections)
	} else {
		posMax = len(appData.StartingStation.Connections)
	}

	visited[start] = struct{}{}
	for _, v := range appData.NetworkMap {

	}
	return res
}
>>>>>>> 504b1aa (algorithm 2 updated)
