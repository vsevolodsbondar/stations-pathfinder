package tests

import (
	"testing"
	c "trains/cli"
	m "trains/models"
	p "trains/pathfinder"
)

func TestShouldFindValidPathsInDenseGraphForAlgorithm2(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/tricky2_test.txt",
		StartingStation: "start",
		EndingStation:   "end",
		TrainNumb:       100,
	}

	appData, err := c.DataConfiguration(config)
	if err != nil {
		t.Fatal(err)
	}

	network := p.BuildFlowGraph(&appData)

	startID := network.StationToID[appData.StartingStation]
	endID := network.StationToID[appData.EndingStation]

	maxFlow := network.Graph.MaxFlow(startID, endID)

	if maxFlow == 0 {
		t.Fatal("expected at least one path")
	}

	paths, err1 := network.ExtractPaths(maxFlow)
	if err1 != nil {
		t.Fatal(err1)
	}

	if len(paths) != maxFlow {
		t.Fatalf("expected %d paths, got %d", maxFlow, len(paths))
	}

	usedStations := make(map[*m.Station]bool)

	for i, path := range paths {

		if len(path) < 2 {
			t.Fatalf("path %d is too short", i)
		}

		if path[0] != appData.StartingStation {
			t.Fatalf("path %d does not start at start station", i)
		}

		if path[len(path)-1] != appData.EndingStation {
			t.Fatalf("path %d does not end at end station", i)
		}

		visited := make(map[*m.Station]bool)

		for j, station := range path {

			// No cycles.
			if visited[station] {
				t.Fatalf("cycle detected in path %d", i)
			}
			visited[station] = true

			// Intermediate stations must be unique across all paths.
			if j != 0 && j != len(path)-1 {
				if usedStations[station] {
					t.Fatalf("station %s belongs to more than one path", station.Name)
				}
				usedStations[station] = true
			}

			// Every consecutive pair must be connected.
			if j < len(path)-1 {
				found := false

				for _, next := range station.Connections {
					if next == path[j+1] {
						found = true
						break
					}
				}

				if !found {
					t.Fatalf(
						"invalid connection %s -> %s",
						station.Name,
						path[j+1].Name,
					)
				}
			}
		}
	}
}
