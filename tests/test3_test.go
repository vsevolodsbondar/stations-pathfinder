package tests

import (
	"fmt"
	"testing"
	c "trains/cli"
	m "trains/models"
	p "trains/pathfinder"
)

func TestShouldFindMoreThanOneRouteFor4Algorithm1(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       4,
	}

	appData, _ := c.DataConfiguration(config)

	routes, _ := p.DFSRangedRouteSets(appData)
	found2Routes := false
	for _, v := range routes {
		for _, r := range v {
			r.PrintRoute()

		}
		if len(v) >= 2 {
			found2Routes = true
		}
	}

	if !found2Routes {
		t.Errorf("Should find at least 2 routes.")
	}

}

func TestShouldFindMoreThanOneRouteFor4Algorithm2(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       4,
	}

	appData, _ := c.DataConfiguration(config)

	res := p.BuildFlowGraph(&appData)
	startID := res.StationToID[appData.StartingStation]
	endID := res.StationToID[appData.EndingStation]

	maxFlow := res.Graph.MaxFlow(startID, endID)
	paths, _ := res.ExtractPaths(maxFlow)

	found2Routes := false
	for _, v := range paths {
		route := ""
		for _, st := range v {
			route += st.Name + " "
		}
		fmt.Println(route)
		if len(v) >= 2 {
			found2Routes = true
		}
	}

	if !found2Routes {
		t.Errorf("Should find at least 2 routes.")
	}
}
