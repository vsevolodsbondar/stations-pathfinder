package validation

import (
	"fmt"
	"strconv"
	"strings"
	m "trains/models"
)

type StartStationValidator struct{}
type EndStationValidator struct{}
type StartEndValidator struct{}
type UniqueCoordinatesForStation struct{}
type StationLineValidator struct{}
type ConnectionLineValidator struct{}

type DuplicateConnectionsSliceValidator struct{}

func (v StartStationValidator) Validate(appData m.AppData) bool {
	// checking some fields to make shure that it was initialized
	if appData.StartingStation.Connections == nil {
		return false
	}
	if !containsStation(appData.NetworkMap, appData.StartingStation.Name) {
		return false
	}

	return true
}

func (v EndStationValidator) Validate(appData m.AppData) bool {
	if appData.EndingStation.Connections == nil {
		return false
	}

	if !containsStation(appData.NetworkMap, appData.EndingStation.Name) {
		return false
	}

	return true
}

func (v DuplicateConnectionsSliceValidator) Validate(connections []string) (bool, error) {
	seen := make(map[string]struct{})

	for _, con := range connections {
		key := normalize(con)

		if _, ok := seen[key]; ok {
			return false, fmt.Errorf("Was seen already: %s", key)
		}

		seen[key] = struct{}{}
	}

	return true, nil
}

func (v UniqueCoordinatesForStation) Validate(appData m.AppData) bool {
	seen := make(map[string]struct{})
	for _, station := range appData.NetworkMap {
		key := strconv.Itoa(station.X_axis) + "," + strconv.Itoa(station.Y_axis)

		if _, ok := seen[key]; ok {
			return false
		}

		seen[key] = struct{}{}
	}

	return true
}

func (v StationLineValidator) Validate(line string) bool {
	line = isComment(line)
	args := strings.Split(line, ",")
	if len(args) != 3 {
		return false
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	if !validName(args[0]) {
		return false
	}
	if !validAxis(args[1]) {
		return false
	}
	if !validAxis(args[2]) {
		println(args[2])
		return false
	}
	return true
}

func (v ConnectionLineValidator) Validate(line string) bool {
	line = isComment(line)
	args := strings.Split(line, "-")
	if len(args) != 2 {
		return false
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	if !validName(args[0]) {
		return false
	}
	if !validName(args[1]) {
		println(args[1])
		return false
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

func validName(name string) bool {
	if name == "" {
		return false
	}
	for _, v := range name {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz_1234567890", v) {
			return false
		}
	}
	return true
}

func validAxis(axis string) bool {
	if axis == "" {
		return false
	}
	for _, v := range axis {
		if !strings.ContainsRune("1234567890", v) {
			return false
		}
	}
	return true
}

func isComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}
