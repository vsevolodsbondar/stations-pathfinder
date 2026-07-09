package io

import (
	"errors"
	"fmt"
	"os"
	"strings"
	s "trains/models"
)

func HandleInitialInputFile(path string) (map[string]s.Station, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	stations := make(map[string]s.Station)

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.New("Cannot read file")
	}

	fmt.Println(string(data))

	return stations, nil

}

func validateStations(line string) bool {
	line = strings.ReplaceAll(line, " ", "")
	args := strings.Split(line, ",")
	if len(args) != 3 {
		return false
	}
	for _, v := range args[0] {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz_", v) {
			return false
		}
	}
	for _, v := range args[1] {
		if !strings.ContainsRune("1234567890", v) {
			return false
		}
	}
	for _, v := range args[2] {
		if !strings.ContainsRune("1234567890", v) {
			return false
		}
	}
	return true
}
