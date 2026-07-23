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

	//res := p.BigFuckingSearch(appData)
	res := p.BuildFlowGraph(&appData)

	maxFlow := res.Graph.MaxFlow(res.StationToID[appData.StartingStation], res.StationToID[appData.EndingStation])
	paths, err := res.ExtractPaths(maxFlow)
	distribution := s.DistributeTrains(paths, appData.TrainNumb)
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

}
