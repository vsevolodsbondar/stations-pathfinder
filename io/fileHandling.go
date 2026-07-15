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

func HandleInitialInputFile(path string) (map[string]*s.Station, error) {
	var conDuplicates v.DuplicateConnectionsSliceValidator
	var stValidator v.StationLineValidator
	var conValidator v.ConnectionLineValidator
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	stations := make(map[string]*s.Station)
	connections := []string{}
	st := false
	con := false

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Cannot open the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineNumb := 0
	for scanner.Scan() {
		lineNumb++
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
				err := WriteStation(stations, line, lineNumb)
				if err != nil {
					return nil, err
				}
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid Station (%s) %d", line, lineNumb)
			}
		case con:
			if conValidator.Validate(line) {
				line = isComment(line)
				connections = append(connections, line)
			} else if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Invalid connection (%s) %d", line, lineNumb)
			}
		default:
			if !(strings.Contains(line, "#") || line == "") {
				return nil, fmt.Errorf("Not commented line (%s) %d", line, lineNumb)
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
	error := WriteConnections(stations, connections)
	if error != nil {
		return nil, error
	}
	return stations, nil
}

func WriteStation(stations map[string]*s.Station, line string, lineNumb int) error {
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
		return fmt.Errorf("X-axis parsing error: %s %d", err, lineNumb)
	}
	y, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("Y-axis parsing error: %s %d", err, lineNumb)
	}
	station := &s.Station{
		Name:   args[0],
		X_axis: x,
		Y_axis: y,
	}
	stations[args[0]] = station

	return nil
}

func WriteConnections(stations map[string]*s.Station, connections []string) error {
	for _, connection := range connections {
		parts := strings.Split(connection, "-")
		if len(parts) != 2 {
			continue
		}

		from := parts[0]
		to := parts[1]

		if from == to {
			return fmt.Errorf("Invalid connection. Can't be same station: %s", connection)
		}

		stations[from].Connections = append(stations[from].Connections, stations[to])
		stations[to].Connections = append(stations[to].Connections, stations[from])
	}

	return nil
}

func isComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}
