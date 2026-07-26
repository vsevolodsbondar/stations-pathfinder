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

func TestShouldScheduleTrainsIn11TurnsForAlgorithm1(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/3_test.txt",
		StartingStation: "beginning",
		EndingStation:   "terminus",
		TrainNumb:       20,
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

	if strings.Contains(res, "T1-terminus") {
		contains1 = true
	}

	if strings.Contains(res, "T20-terminus") {
		contains2 = true
	}

	if !contains1 && !contains2 {
		t.Errorf("Should contain at least T1-terminus and T20-terminus in output.")
	}

	moveCounter := 0
	lines := strings.Split(res, "\n")
	for _, v := range lines {
		strings.TrimSpace(v)
		if v != "" {
			moveCounter++
		}
	}

	if moveCounter > 11 {
		t.Errorf("Should end scheduling in 11 turns.")
	}
}

func TestShouldScheduleTrainsIn11TurnsForAlgorithm2(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/3_test.txt",
		StartingStation: "beginning",
		EndingStation:   "terminus",
		TrainNumb:       20,
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
		if strings.Contains(line, "T1-terminus") {
			contains1 = true
		}

		if strings.Contains(line, "T20-terminus") {
			contains2 = true
		}

		fmt.Println(line)
	}

	if !contains1 && !contains2 {
		t.Errorf("Should contain at least T1-terminus and T20-terminus in output.")
	}

	if len(lines) > 11 {
		t.Errorf("Should end scheduling in 11 turns.")
	}
}
