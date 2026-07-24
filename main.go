package main

import (
	"fmt"
	"os"
	c "trains/cli"
	p "trains/pathfinder"
	s "trains/scheduler"
)

func main() {
	//go run . -feature stations.map waterloo euston 5
	//go run . -feature jungle-desert.map jungle desert 5
	//go run . -feature smallAndLarge.map small large 9
	conf, err := c.FlagHandling()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	appData, errs := c.DataConfiguration(conf)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
		os.Exit(1)
	}

	res := p.BuildFlowGraph(&appData)
	startID := res.StationToID[appData.StartingStation]
	endID := res.StationToID[appData.EndingStation]
	trains := appData.TrainNumb

	maxFlow := res.Graph.MaxFlow(startID, endID)
	paths, err := res.ExtractPaths(maxFlow)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	distribution := s.DistributeTrains(paths, trains)
	for _, v := range distribution {
		fmt.Println(v)
	}

	for i, path := range paths {
		fmt.Printf("%d: ", i)
		for _, st := range path {
			fmt.Printf("%s ", st.Name)
		}
		fmt.Println()
	}

	lines := s.Schedule(paths, trains)

	for _, line := range lines {
		fmt.Println(line)
	}

}
