package main

import (
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

	res.Graph.Print()
}
