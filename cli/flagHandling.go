package cli

import (
	"flag"
	"fmt"
	"strconv"
	m "trains/models"
)

func FlagHandling() (m.FlagConfig, error) {
	someFeature := flag.Bool("feature", false, "does additional stuff during the run.") //using it as a placeholder for some features to be activated with a flag
	flag.Parse()

	args := flag.Args()

	if len(args) != 4 {
		err := fmt.Errorf("Usage: [flags] <network file> <start station> <end station> <train amount>\n")
		return m.FlagConfig{}, err
	}

	networkFile := args[0]
	start := args[1]
	end := args[2]

	if start == end {
		err := fmt.Errorf("Starting and ending locations should not be the same")
		return m.FlagConfig{}, err
	}

	trains, err := strconv.Atoi(args[3])
	if err != nil || trains <= 0 {
		err := fmt.Errorf("Trains must be a positive integer")
		return m.FlagConfig{}, err
	}

	return m.FlagConfig{
		NetworkMapPath:  networkFile,
		StartingStation: start,
		EndingStation:   end,
		TrainNumb:       trains,
		Feature:         *someFeature,
	}, nil
}
