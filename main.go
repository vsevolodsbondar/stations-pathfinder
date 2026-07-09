package main

import (
	"fmt"
	"log"
	c "trains/cli"
	io "trains/io"
)

func main() {
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}
	stations, err := (io.HandleInitialInputFile("stations.map"))
	if err != nil {
		log.Fatal()
	}
	for _, station := range stations {
		fmt.Print(station.Name + ": ")
		for _, con := range station.Connections {
			fmt.Println(con.Name)
		}
		println()
	}
	fmt.Println(conf.NetworkMapPath)
	fmt.Println(conf.StartingStation)
	fmt.Println(conf.EndingStation)
	fmt.Println(conf.TrainNumb)
}
