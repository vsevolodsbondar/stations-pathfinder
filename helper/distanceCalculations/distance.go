package distancecalculations

import (
	"math"

	m "trains/models"
)

func PathsDistance(appData m.AppData, path []string) m.Route {
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
