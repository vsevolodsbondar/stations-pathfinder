package main

import (
	"fmt"
	"os"
	c "trains/cli"
	p "trains/pathfinder"
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

	res, err := p.DFSRangedRouteSets(appData)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	for i, v := range res {
		fmt.Println("Set", i+1, ":")
		for _, r := range v {
			r.PrintRoute()
		}
	}
}
