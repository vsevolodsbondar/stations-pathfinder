package validation

import (
	"fmt"
	"os"
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
type StationConnectionBlocks struct{}

type DuplicateConnectionsSliceValidator struct{}

func (v StationConnectionBlocks) Validate(path string) (bool, []error) {
	file, err := os.ReadFile(path)
	errors := []error{}
	if err != nil {
		errors = append(errors, err)
	}
	stations, _, ok := strings.Cut(string(file), "connections")
	if !ok {
		errors = append(errors, fmt.Errorf("No connections block in file"))
	}
	_, stations, ok = strings.Cut(stations, "stations")
	if !ok {
		errors = append(errors, fmt.Errorf("No stations block in file"))
	}
	lines := strings.Split(stations, "\n")
	count := 0
	for _, line := range lines {
		isComment(line)
		if line != "" {
			count++
		}
		if count > 10000 {
			errors = append(errors, fmt.Errorf("map contains more than 10000 stations"))
			break
		}
	}
	return true, errors
}

func (v StartStationValidator) Validate(appData m.AppData) (bool, []error) {
	valid := true
	errs := []error{}

	// checking some fields to make shure that it was initialized
	if appData.StartingStation == nil {
		valid = false
		errs = append(errs, fmt.Errorf("Starting station does not exist."))
		return valid, errs
	}

	if appData.StartingStation.Connections == nil {
		valid = false
		errs = append(errs, fmt.Errorf("Starting station doesn't connect to anything."))
	}

	if !containsStation(appData.NetworkMap, appData.StartingStation.Name) {
		valid = false
		errs = append(errs, fmt.Errorf("Starting station is not in the graph."))
	}

	return valid, errs
}

func (v EndStationValidator) Validate(appData m.AppData) (bool, []error) {
	valid := true
	errs := []error{}

	if appData.EndingStation == nil {
		valid = false
		errs = append(errs, fmt.Errorf("Ending station does not exist."))
		return valid, errs
	}

	if appData.EndingStation.Connections == nil {
		valid = false
		errs = append(errs, fmt.Errorf("Ending station doesn't connect to anything."))
	}

	if !containsStation(appData.NetworkMap, appData.EndingStation.Name) {
		valid = false
		errs = append(errs, fmt.Errorf("Ending station is not in the graph."))
	}

	return valid, errs
}

func (v DuplicateConnectionsSliceValidator) Validate(connections []string) (bool, []error) {
	valid := true
	errs := []error{}

	seen := make(map[string]struct{})

	for _, con := range connections {
		key := normalize(con)

		if _, ok := seen[key]; ok {
			valid = false
			errs = append(errs, fmt.Errorf("Was seen already: %s", key))
		}

		seen[key] = struct{}{}
	}

	return valid, errs
}

func (v UniqueCoordinatesForStation) Validate(appData m.AppData) (bool, []error) {
	valid := true
	errs := []error{}

	seen := make(map[string]string)
	for _, station := range appData.NetworkMap {
		key := strconv.Itoa(station.X_axis) + "," + strconv.Itoa(station.Y_axis)

		if first, ok := seen[key]; ok {
			valid = false
			errs = append(errs,
				fmt.Errorf("duplicate coordinates (%d, %d): stations %q and %q",
					station.X_axis,
					station.Y_axis,
					first,
					station.Name))
		}

		seen[key] = station.Name
	}

	return valid, errs
}

func (v StationLineValidator) Validate(line string) (bool, []error) {
	errs := []error{}
	valid := true

	line = isComment(line)
	args := strings.Split(line, ",")
	if len(args) != 3 {
		valid = false
		errs = append(errs, fmt.Errorf("Not valid station line: %s. Station line doesn't contain 3 elements.", line))
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	if len(args) == 3 {
		ok, err := validName(args[0])
		if !ok {
			valid = false
			errs = append(errs, err)
		}

		ok, err = validAxis(args[1])
		if !ok {
			valid = false
			errs = append(errs, err)
		}

		ok, err = validAxis(args[2])
		if !ok {
			valid = false
			errs = append(errs, err)
		}
	}

	return valid, errs
}

func (v ConnectionLineValidator) Validate(line string) (bool, []error) {
	errs := []error{}
	valid := true

	line = isComment(line)
	args := strings.Split(line, "-")
	if len(args) != 2 {
		valid = false
		errs = append(errs, fmt.Errorf("Not a valid connection line: %s. Connection contains not 2 stations.", line))
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	if len(args) == 2 {
		_, err := validName(args[0])
		if err != nil {
			valid = false
			errs = append(errs, err)
		}

		_, error := validName(args[1])
		if error != nil {
			valid = false
			errs = append(errs, err)
		}
	}

	return valid, errs
}

func containsStation(station []*m.Station, name string) bool {
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

func validName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("Name is empty.")
	}
	for _, v := range name {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz_1234567890", v) {
			return false, fmt.Errorf("Not valid symbol in name: %s.", string(v))
		}
	}
	return true, nil
}

func validAxis(axis string) (bool, error) {
	if axis == "" {
		return false, fmt.Errorf("Axis is empty.")
	}
	for _, v := range axis {
		if !strings.ContainsRune("1234567890", v) {
			return false, fmt.Errorf("Axis contains not valid symbol: %s.", string(v))
		}
	}
	return true, nil
}

func isComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}
