package main

import (
	"fmt"
	"log"
	c "trains/cli"
)

func main() {
	//go run . -feature stations.map waterloo euston 5
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}

	appData, err := c.DataConfiguration(conf)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range appData.NetworkMap {
		fmt.Println("Station:", v.Name)
		for _, con := range v.Connections {
			fmt.Println("Connection:", con.Name)
		}
	}
}
