package io

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	s "trains/models"
)

func HandleInitialInputFile(path string) (map[string]s.Station, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	stations := make(map[string]s.Station)

	file, err := os.Open("stations.map")
	if err != nil {
		return nil, errors.New("Cannot open the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stations, nil

}

func validateStations(line string) bool {
	line = strings.ReplaceAll(line, " ", "")
	args := strings.Split(line, ",")
	if len(args) != 3 {
		return false
	}
	for _, v := range args[0] {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz_1234567890", v) {
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
