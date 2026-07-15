package main

import (
	"fmt"
	"log"
	c "trains/cli"
	p "trains/pathfinder"
)

func main() {
	//go run . -feature stations.map waterloo euston 5
	//go run . -feature test2.map jungle desert 5
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	appData, err := c.DataConfiguration(conf)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	for _, v := range appData.NetworkMap {
		fmt.Println("Station:", v.Name)
		for _, con := range v.Connections {
			fmt.Println("Connection:", con.Name)
		}
	}

	res := p.DFSRangeRoutes(appData)
	for _, v := range res {
		fmt.Println(v.Route)
		fmt.Println(v.Distance)
	}
}
