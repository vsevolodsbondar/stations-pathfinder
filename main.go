package main

import (
	"fmt"
	"log"
	"os"
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

	appData, errs := c.DataConfiguration(conf)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Println(err)
		}
		os.Exit(1)
	}

	for _, v := range appData.NetworkMap {
		fmt.Println("Station:", v.Name)
		for _, con := range v.Connections {
			fmt.Println("Connection:", con.Name)
		}
	}

	res := p.BigFuckingSearch(appData)
	for _, v := range res {
		fmt.Println(v)
	}
}
