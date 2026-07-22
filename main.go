package main

import (
	"fmt"
	"log"
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
		log.Fatal("Error: ", err)
	}

	appData, err := c.DataConfiguration(conf)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	res, err := p.DFSRangedRouteSets(appData)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	shortest := s.BestSingleRoute(res)
	fmt.Println("Shortest:")
	shortest[0].PrintRoute()

	multiple := s.BestMultipleRoutes(res, appData.TrainNumb)
	fmt.Println("For trains:", appData.TrainNumb)
	for _, v := range multiple {
		v.PrintRoute()
	}

	s.MoveTrains(multiple, appData.TrainNumb)
}
