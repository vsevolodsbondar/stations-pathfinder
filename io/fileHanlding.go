package io

import (
	"bufio"
	"errors"
	"fmt"
	"log"
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
				fmt.Println(strings.Contains(line, "#"))
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

func WriteStation(stations map[string]s.Station, line string) error {
	var station s.Station
	st, comment, ok := strings.Cut(line, "#")
	if ok {
		line = st
		log.Print(comment)
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

func isComment(line string) string {
	if strings.Contains(line, "#") {
		line, _, _ = strings.Cut(line, "#")
	}
	return line
}
