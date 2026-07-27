package tests

import (
	"testing"
	c "trains/cli"
	m "trains/models"
	p "trains/pathfinder"
)

func TestShouldFind2PathsWithManyAlternativeRoutes(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/tricky1_test.txt",
		StartingStation: "start",
		EndingStation:   "end",
		TrainNumb:       10000,
	}

	appData, err1 := c.DataConfiguration(config)
	if len(err1) != 0 {
		t.Fatalf("configuration failed: %v", err1)
	}

	network := p.BuildFlowGraph(&appData)

	startID := network.StationToID[appData.StartingStation]
	endID := network.StationToID[appData.EndingStation]

	maxFlow := network.Graph.MaxFlow(startID, endID)

	if maxFlow != 2 {
		t.Fatalf("expected max flow 2, got %d", maxFlow)
	}

	paths, err2 := network.ExtractPaths(maxFlow)
	if err2 != nil {
		t.Fatalf("failed to extract paths: %v", err2)
	}

	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(paths))
	}

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
	}
}
