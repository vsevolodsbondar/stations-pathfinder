package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	v "trains/helper/validation"
	s "trains/models"
)

func HandleInitialInputFile(path string) (map[string]s.Station, error) {
	var conDuplicates v.DuplicateConnectionsSliceValidator
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	stations := make(map[string]s.Station)
	connections := []string{}
	st := false
	con := false

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Cannot open the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if strings.Contains(line, "stations:") {
			st = true
			continue
		}
		if strings.Contains(line, "connections:") {
			st = false
			con = true
			continue
		}

		if st {
			if ValidateStations(line) {
				WriteStation(stations, line)
			} else {
				return nil, fmt.Errorf("Invalid Station (%s)", line)
			}
		}
		if con {
			if ValidateConnections(line) {
				line = IsComment(line)
				connections = append(connections, line)
			} else {
				return nil, fmt.Errorf("Invalid connection (%s)", line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error reading map files")
	}
	ok, err := conDuplicates.Validate(connections)
	if !ok {
		return nil, err
	}
	WriteConnections(stations, connections)
	return stations, nil
}

func ValidateStations(line string) bool {
	line = IsComment(line)
	args := strings.Split(line, ",")
	if len(args) != 3 {
		return false
	}
	if !ValidName(args[0]) {
		return false
	}
	if !ValidAxis(args[1]) {
		return false
	}
	if !ValidAxis(args[2]) {
		return false
	}
	return true
}

func ValidateConnections(line string) bool {
	line = IsComment(line)
	args := strings.Split(line, "-")
	if len(args) != 2 {
		return false
	}
	if !ValidName(args[0]) {
		return false
	}
	if !ValidName(args[1]) {
		return false
	}
	return true
}

func WriteStation(stations map[string]s.Station, line string) {
	var station s.Station
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	args := strings.Split(line, ",")
	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[1])
	station.Name = args[0]
	station.X_axis = x
	station.Y_axis = y
	stations[args[0]] = station
}

func IsComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}

func WriteConnections(stations map[string]s.Station, connections []string) {
	for _, v := range connections {
		for key := range stations {
			if strings.Contains(v, key) {
				name := strings.ReplaceAll(v, key, "")
				name = strings.ReplaceAll(name, "-", "")
				if station, ok := stations[key]; ok {
					station.Connections = append(station.Connections, stations[name])
					stations[key] = station
				}
			}
		}
	}
}

func ValidName(name string) bool {
	for _, v := range name {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz_1234567890", v) {
			return false
		}
	}
	return true
}

func ValidAxis(axis string) bool {
	for _, v := range axis {
		if !strings.ContainsRune("1234567890", v) {
			return false
		}
	}
	return true
}
