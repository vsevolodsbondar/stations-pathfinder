package tests

import (
	"fmt"
	"testing"
	c "trains/cli"
	r "trains/helper/routeUtils"
	m "trains/models"
	p "trains/pathfinder"
)

func TestShouldFindOneRouteFor1Algorithm1(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       1,
	}

	appData, _ := c.DataConfiguration(config)

	routes, _ := p.DFSRangedRouteSets(appData)
	best := r.BestRoutes(routes, appData.TrainNumb)

	for _, r := range best {
		r.PrintRoute()
	}

	if len(best) > 1 {
		t.Errorf("Should find find exactly 1 route.")
	}
}

func TestShouldFindOneRouteFor1Algorithm2(t *testing.T) {
	config := m.FlagConfig{
		NetworkMapPath:  "./testData/waterloo.txt",
		StartingStation: "waterloo",
		EndingStation:   "st_pancras",
		TrainNumb:       1,
	}

	appData, _ := c.DataConfiguration(config)

	res := p.BuildFlowGraph(&appData)
	startID := res.StationToID[appData.StartingStation]
	endID := res.StationToID[appData.EndingStation]

	maxFlow := res.Graph.MaxFlow(startID, endID)
	paths, _ := res.ExtractPaths(maxFlow)

	found1Route := false
	for _, v := range paths {
		route := ""
		for _, st := range v {
			route += st.Name + " "
		}
		fmt.Println(route)
		if len(v) >= 2 {
			found1Route = true
		}
	}

	if !found1Route {
		t.Errorf("Should find at least 2 routes.")
	}
}
