package main

import (
	"fmt"
	"log"
	c "trains/cli"
	p "trains/pathfinder"
)

func main() {
	//go run . -feature stations.map waterloo euston 5
	//go run . -feature jungle-desert.map jungle desert 5
	//go run . -feature smallAndLarge.map small large 9
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	appData, err := c.DataConfiguration(conf)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	//res := p.BigFuckingSearch(appData)
	res := p.BuildFlowGraph(&appData)

	maxFlow := res.Graph.MaxFlow(res.StationToID[appData.StartingStation], res.StationToID[appData.EndingStation])
	paths, err := res.ExtractPaths(maxFlow)

	for i, path := range paths {
		fmt.Printf("%d: ", i)
		for _, st := range path {
			fmt.Printf("%s ", st.Name)
		}
		fmt.Println()
	}

}
