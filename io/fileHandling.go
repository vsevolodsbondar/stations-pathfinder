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

func HandleInitialInputFile(path string) (map[string]*s.Station, []error) {
	var conDuplicates v.DuplicateConnectionsSliceValidator
	var stValidator v.StationLineValidator
	var conValidator v.ConnectionLineValidator
	var stConBlocks v.StationConnectionBlocks
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, []error{errors.New("File does not exist")}
	}
	stations := make(map[string]*s.Station)
	connections := []string{}
	st := false
	con := false
	errs := []error{}

	val, valErrs := stConBlocks.Validate(path)
	if !val {
		errs = append(errs, valErrs...)
		return nil, errs
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, []error{errors.New("Cannot open the file")}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineNumb := 0
	for scanner.Scan() {
		lineNumb++
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")

		//skipping empty lines
		line = isComment(line)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

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
			ok, validationErrs := stValidator.Validate(line)
			if validationErrs != nil {
				for _, err := range validationErrs {
					errs = append(errs,
						fmt.Errorf("Invalid station (%s), line %d: %w", line, lineNumb, err))
				}
			}
			if ok {
				err := WriteStation(stations, line, lineNumb)
				if err != nil {
					errs = append(errs, err)
				}
			}
		case con:
			ok, validationErrs := conValidator.Validate(line)
			if validationErrs != nil {
				for _, err := range validationErrs {
					errs = append(errs,
						fmt.Errorf("Invalid connection: %s, line: %d: %w", line, lineNumb, err))
				}
			}
			if ok {
				line = isComment(line)
				connections = append(connections, line)
			}
		default:
			if !(strings.Contains(line, "#") || line == "") {
				errs = append(errs, fmt.Errorf("Not commented line: %s, line: %d", line, lineNumb))
			}
		}

	}

	if len(errs) > 0 {
		return nil, errs
	}

	if err := scanner.Err(); err != nil {
		return nil, []error{errors.New("Error reading map files")}
	}
	ok, errs := conDuplicates.Validate(connections)
	if !ok {
		return nil, errs
	}
	err = WriteConnections(stations, connections)
	if err != nil {
		return nil, []error{err}
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

	stationInMap, ok := stations[args[0]]
	if ok {
		return fmt.Errorf("Station is duplicated: %s, line: %d", stationInMap.Name, lineNumb)
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
