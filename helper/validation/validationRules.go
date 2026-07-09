package validation

import (
	"strconv"
	"strings"
	m "trains/models"
)

type StartStationValidator struct{}
type EndStationValidator struct{}
type StartEndValidator struct{}
type DuplicateConnectionsValidator struct{} //test if there is same and if there is the same in reverse
type DuplicateConnectionsMapValidator struct{}
type UniqueCoordinatesForStation struct{}

type DuplicateConnectionsSliceValidator struct{}

func (v StartStationValidator) Validate(appData m.AppData) bool {
	// checking some fields to make shure that it was initialized
	if appData.StartingStation.Connections == nil || appData.StartingStation.Name == "" {
		return false
	}
	if !containsStation(appData.NetworkMap, appData.StartingStation.Name) {
		return false
	}

	return true
}

func (v EndStationValidator) Validate(appData m.AppData) bool {
	if appData.EndingStation.Connections == nil || appData.EndingStation.Name == "" {
		return false
	}

	if !containsStation(appData.NetworkMap, appData.EndingStation.Name) {
		return false
	}

	return true
}

func (v StartEndValidator) Validate(appData m.AppData) bool {
	if appData.StartingStation.Name == appData.EndingStation.Name {
		return false
	}

	if appData.StartingStation.X_axis == appData.EndingStation.X_axis &&
		appData.StartingStation.Y_axis == appData.EndingStation.Y_axis {
		return false
	}

	return true
}

func (v DuplicateConnectionsValidator) Validate(appData m.AppData) bool {
	//iterating through all Stations in app data
	for _, station := range appData.NetworkMap {
		connection := ""
		connectionReversed := ""
		//iterating through all Station's connections on by one
		//taking first one and starting to compare to others
		for i, con := range station.Connections {
			connection = con.Name
			parts := strings.Split(con.Name, "-")
			connectionReversed = parts[1] + "-" + parts[0]

			for j, inner := range station.Connections {
				if i == j {
					continue
				}
				if inner.Name == connection || inner.Name == connectionReversed {
					return false
				}
			}
		}
	}

	return true
}

func (v DuplicateConnectionsMapValidator) Validate(appData m.AppData) bool {
	for _, station := range appData.NetworkMap {
		seen := make(map[string]struct{})

		for _, con := range station.Connections {
			key := normalize(con.Name)

			if _, ok := seen[key]; ok {
				return false
			}

			seen[key] = struct{}{}
		}
	}

	return true
}

func (v DuplicateConnectionsSliceValidator) Validate(connections []string) bool {
	seen := make(map[string]struct{})

	for _, con := range connections {
		key := normalize(con)

		if _, ok := seen[key]; ok {
			return false
		}

		seen[key] = struct{}{}
	}

	return true
}

func (v UniqueCoordinatesForStation) Validate(appData m.AppData) bool {
	for _, station := range appData.NetworkMap {
		seen := make(map[string]struct{})

		for _, con := range station.Connections {
			key := strconv.Itoa(con.X_axis) + strconv.Itoa(con.Y_axis)

			if _, ok := seen[key]; ok {
				return false
			}

			seen[key] = struct{}{}
		}
	}

	return true
}

func containsStation(station []m.Station, name string) bool {
	for _, v := range station {
		if v.Name == name {
			return true
		}
	}

	return false
}

// normalize "B-A" to "A-B"
// assuming I always have valid connection
func normalize(name string) string {
	parts := strings.Split(name, "-")

	if parts[0] < parts[1] {
		return parts[0] + "-" + parts[1]
	}

	return parts[1] + "-" + parts[0]
}
