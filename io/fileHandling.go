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
	var stValidator v.StationLineValidator
	var conValidator v.ConnectionLineValidator
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
			con = false
			continue
		}
		if strings.Contains(line, "connections:") {
			st = false
			con = true
			continue
		}

		switch {
		case st:
			if stValidator.Validate(line) {
				err := WriteStation(stations, line)
				if err != nil {
					return nil, err
				}
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid Station (%s)", line)
			}
		case con:
			if conValidator.Validate(line) {
				line = isComment(line)
				connections = append(connections, line)
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid connection (%s)", line)
			}
		default:
			if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Not commented line (%s)", line)
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

func HandleInitialInputFileWithPointer(path string) (map[string]*s.PointerStation, error) {
	var conDuplicates v.DuplicateConnectionsSliceValidator
	var stValidator v.StationLineValidator
	var conValidator v.ConnectionLineValidator
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	stations := make(map[string]*s.PointerStation)
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
			con = false
			continue
		}
		if strings.Contains(line, "connections:") {
			st = false
			con = true
			continue
		}

		switch {
		case st:
			if stValidator.Validate(line) {
				err := WritePointerStation(stations, line)
				if err != nil {
					return nil, err
				}
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid Station (%s)", line)
			}
		case con:
			if conValidator.Validate(line) {
				line = isComment(line)
				connections = append(connections, line)
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid connection (%s)", line)
			}
		default:
			if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Not commented line (%s)", line)
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
	WritePointerConnections(stations, connections)
	return stations, nil
}

func WriteStation(stations map[string]s.Station, line string) error {
	var station s.Station
	st, _, ok := strings.Cut(line, "#")
	if ok {
		line = st
	}
	args := strings.Split(line, ",")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	x, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("X-axis parsing error: %s", err)
	}
	y, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("Y-axis parsing error: %s", err)
	}
	station.Name = args[0]
	station.X_axis = x
	station.Y_axis = y
	stations[args[0]] = station

	return nil
}

func WritePointerStation(stations map[string]*s.PointerStation, line string) error {
	st, _, ok := strings.Cut(line, "#")
	if ok {
		line = st
	}
	args := strings.Split(line, ",")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	x, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("X-axis parsing error: %s", err)
	}
	y, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("Y-axis parsing error: %s", err)
	}
	station := &s.PointerStation{
		Name:   args[0],
		X_axis: x,
		Y_axis: y,
	}
	stations[args[0]] = station

	return nil
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

func WritePointerConnections(stations map[string]*s.PointerStation, connections []string) {
	for _, connection := range connections {
		parts := strings.Split(connection, "-")
		if len(parts) != 2 {
			continue
		}

		from := parts[0]
		to := parts[1]

		stations[from].Connections = append(stations[from].Connections, stations[to])
		stations[to].Connections = append(stations[to].Connections, stations[from])
	}
}

func isComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}
