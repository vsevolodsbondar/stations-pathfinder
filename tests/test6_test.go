package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"
	c "trains/cli"
	m "trains/models"
	p "trains/pathfinder"
	sc "trains/scheduler"
	s "trains/service"
)

func TestShouldPrintTrainsInExpectedFormatForAlgorithm1(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       3,
	}

	ctx := context.Background()

	appData, _ := c.DataConfiguration(config)

	res, err := CaptureStdout(func() error {
		return s.Algorithm1Runner(ctx, appData)
	})

	// fmt.Println("Captured:", res)

	if err != nil {
		t.Fatal(err)
	}

	contains1 := false
	contains2 := false

	if strings.Contains(res, "T1-waterloo") {
		contains1 = true
	}

	if strings.Contains(res, "T1-st_pancras") {
		contains2 = true
	}

	if !contains1 && !contains2 {
		t.Errorf("Should contain at least T1-waterloo and T1-st_pancras in output.")
	}
}

func TestShouldPrintTrainsInExpectedFormatForAlgorithm2(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       3,
	}

	appData, _ := c.DataConfiguration(config)

	res := p.BuildFlowGraph(&appData)
	startID := res.StationToID[appData.StartingStation]
	endID := res.StationToID[appData.EndingStation]

	maxFlow := res.Graph.MaxFlow(startID, endID)
	paths, _ := res.ExtractPaths(maxFlow)

	sc.DistributeTrains(paths, appData.TrainNumb)
	lines := sc.Schedule(paths, appData.TrainNumb)

	contains1 := false
	contains2 := false

	for _, line := range lines {
		if strings.Contains(line, "T1-waterloo") {
			contains1 = true
		}

		if strings.Contains(line, "T1-st_pancras") {
			contains2 = true
		}

		fmt.Println(line)
	}

	if !contains1 && !contains2 {
		t.Errorf("Should contain at least T1-waterloo and T1-st_pancras in output.")
	}
}
