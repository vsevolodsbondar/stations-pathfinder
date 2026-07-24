package cli

import (
	"flag"
	"fmt"
	"strconv"
	m "trains/models"
)

func FlagHandling() (m.FlagConfig, error) {
	algorithm := flag.String("algorithm", "", "switches to Seva's or Anatolii's algorithm for pathfinding and scheduling. To use it: --algorithm=<seva or anatolii>")
	flag.Parse()

	args := flag.Args()

	if len(args) != 4 {
		return m.FlagConfig{}, fmt.Errorf("Usage: --algorithm=<seva or anatolii> <network file> <start station> <end station> <train amount>\n")
	}

	networkFile := args[0]
	start := args[1]
	end := args[2]

	if start == end {
		return m.FlagConfig{}, fmt.Errorf("Starting and ending locations should not be the same")
	}

	trains, err := strconv.Atoi(args[3])
	if err != nil || trains <= 0 {
		return m.FlagConfig{}, fmt.Errorf("Trains must be a positive integer")
	}

	if trains > 20000 {
		return m.FlagConfig{}, fmt.Errorf("Not a valid number of trains. Why do we need that much?:). Max: 20000")
	}

	if *algorithm != "seva" && *algorithm != "anatolii" {
		return m.FlagConfig{}, fmt.Errorf("Unknown algorithm to run.")
	}

	return m.FlagConfig{
		NetworkMapPath:  networkFile,
		StartingStation: start,
		EndingStation:   end,
		TrainNumb:       trains,
		Algorithm:       *algorithm,
	}, nil
}
